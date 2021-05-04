package server

import (
	"context"
	"fmt"
	"net/http"
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
	e := gin.New()
	e.Use(gin.Recovery(), gin.Logger(), cors.Default())
	setRouter(engine, e)
	gin.SetMode(gin.ReleaseMode)

	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	}

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
func setRouter(e *engine.Engine, r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "pong"})
	})

	r.GET("/gen", passwordGenHandler(e))

	r.Static("/assets", "./templates/assets")
	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
}

// Start initializes the Server.
func (s *Server) Start() error {
	if err := s.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("initiating server: %w", err)
	}

	return nil
}

// Shutdown gracefully shutdowns the Server. During the given
// duration time the Server won't receive any new requests but it will
// finish all pending processes.
func (s *Server) Shutdown(duration time.Duration) (err error) {
	defer func() {
		err = s.srv.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	if err = s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("graceful shutdown (duration: %v): %w", duration, err)
	}

	return
}
