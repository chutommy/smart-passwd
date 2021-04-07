package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

// Configuration file default values.
var (
	FileName = "config"
	FileType = "yaml"
	FilePath = "."
)

// Key values of the configuration.
var (
	KeyHTTPPort = "HTTPPort"
	KeyDBFile   = "DBFile"
	KeyDebug    = "Debug"
)

// GetConfig sets defaults, replaces them with the values from the configuration file
// and finally overrides them with flags.
func GetConfig(args []string) (*viper.Viper, error) {
	vi := viper.New()

	// default
	vi.SetDefault(KeyHTTPPort, 8080)
	vi.SetDefault(KeyDBFile, "./words.db")
	vi.SetDefault(KeyDebug, false)

	vi.SetConfigName(FileName)
	vi.SetConfigType(FileType)
	vi.AddConfigPath(FilePath)

	// config file
	err := vi.ReadInConfig()
	if err != nil {
		if !errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("config file: %w", err)
		}

		err = viper.SafeWriteConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to write config file: %w", err)
		}
	}

	// flags
	fs := flag.NewFlagSet("Smart Pass-WD", flag.ExitOnError)

	http := fs.Int64("http", vi.GetInt64(KeyHTTPPort), "port of the application to serve")
	db := fs.String("db", vi.GetString(KeyDBFile), "path to SQLite3 database file")
	debug := fs.Bool("debug", vi.GetBool(KeyDebug), "debug mode")

	err = fs.Parse(args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}

	vi.Set(KeyHTTPPort, *http)
	vi.Set(KeyDBFile, *db)
	vi.Set(KeyDebug, *debug)

	return vi, nil
}
