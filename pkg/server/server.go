package server

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/chutified/smart-passwd/pkg/config"
	"github.com/chutified/smart-passwd/pkg/engine"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server represents a server node of the web service.
type Server struct {
	engine *engine.Engine
	srv    *http.Server
}

// NewServer constructs a new Server from the given dependent modules.
func NewServer(cfg *config.Config, engine *engine.Engine) *Server {
	gin.SetMode(gin.ReleaseMode)

	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	}

	e := gin.New()
	e.Use(gin.Recovery(), gin.Logger(), cors.Default())
	setRouter(cfg.RootPath, engine, e)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler:           e,
		ReadTimeout:       1000 * time.Millisecond,
		ReadHeaderTimeout: 500 * time.Millisecond,
		WriteTimeout:      1000 * time.Millisecond,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	return &Server{
		engine: engine,
		srv:    srv,
	}
}

// setRouter sets routes for the engine (gin).
func setRouter(root string, e *engine.Engine, r *gin.Engine) {
	r.GET("/ping", pingHandler())
	r.POST("/gen", passwordGenHandler(e))

	r.Static("/assets", filepath.Join(root, "public/assets"))
	r.Static("/scripts", filepath.Join(root, "public/scripts"))
	r.LoadHTMLFiles(filepath.Join(root, "public/index.html"))
	r.GET("/", homePageHandler())
}

// Start initializes the Server.
func (s *Server) Start() error {
	return fmt.Errorf("server listens and servers: %w", s.srv.ListenAndServe())
}

// Shutdown gracefully shutdowns the Server. During the given
// duration time the Server won't receive any new requests but it will
// finish all pending processes.
func (s *Server) Shutdown(duration time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	if err = s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("graceful shutdown (duration: %v): %w", duration, err)
	}

	return
}
