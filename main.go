package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/chutified/smart-passwd/pkg/config"
	"github.com/chutified/smart-passwd/pkg/data"
	"github.com/chutified/smart-passwd/pkg/engine"
	"github.com/chutified/smart-passwd/pkg/server"
	"github.com/chutified/smart-passwd/pkg/utils"
	_ "modernc.org/sqlite"
)

func main() {
	defer os.Exit(0)

	logger := log.Default()
	logger.Printf("setting configuration...\n")

	cfg, err := loadConfig()
	if err != nil {
		logger.Printf("failed to retrieve configuration: %v\n", err)
		runtime.Goexit()
	}

	logger.Printf("configuration successfully retrieved\n")
	logger.Printf("connecting to a database...\n")

	wl, err := connectDB(cfg)
	if err != nil {
		logger.Printf("failed to connect to the database: %v\n", err)
		runtime.Goexit()
	}

	defer func() {
		if err := wl.Close(); err != nil {
			logger.Printf("failed to close the database connection: %v\n", err)
			runtime.Goexit()
		}
	}()

	logger.Printf("successfully connected to the database: %s\n", cfg.DBFile)
	logger.Printf("setting the server...\n")

	srv := setServer(wl, cfg)

	logger.Printf("server successfully set\n")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		logger.Printf("terminate signal received: %v\n", <-c)

		if err = srv.Shutdown(10 * time.Second); err != nil {
			logger.Printf("failed to gracefully shutdown: %v\n", err)
			runtime.Goexit()
		}
	}()

	logger.Printf("launching a http server...\n")
	logger.Printf("open in browser: http://127.0.0.1:%d\n", cfg.HTTPPort)

	if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Printf("failed to start the server: %v\n", err)
		runtime.Goexit()
	}

	logger.Printf("server successfully closed\n")
}

func setServer(wl *data.SQLiteWordList, cfg *config.Config) *server.Server {
	ctr := engine.NewConstructor(3, 22)
	swp := engine.NewSwapper()
	e := engine.Init(wl, ctr, swp)

	return server.NewServer(cfg, e)
}

func connectDB(cfg *config.Config) (*data.SQLiteWordList, error) {
	dbDir, dbBase := filepath.Split(cfg.DBFile)
	dbFileArr := strings.Split(dbBase, ".")
	dbFile := utils.NewFile(dbDir, dbFileArr[0], dbFileArr[1])

	wl, err := data.ConnectSQLite(dbFile)
	if err != nil {
		return nil, fmt.Errorf("data connection: %w", err)
	}

	return wl, nil
}

func loadConfig() (*config.Config, error) {
	defaultCfg := config.NewConfig(8080, ".", true, ".")
	fileCfg := utils.NewFile(".", "config", "yaml")

	cfg, err := config.GetConfig(defaultCfg, fileCfg, os.Args)
	if err != nil {
		return nil, fmt.Errorf("get config: %w", err)
	}

	return cfg, nil
}
