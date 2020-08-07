package server

import (
	"fmt"
	"net/http"
	"time"

	config "github.com/chutified/smart-passwd/config"
	handlers "github.com/chutified/smart-passwd/handlers"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Server defines the server settings.
type Server struct {
	srv *http.Server
	ph  *handlers.PWDhandler
}

// New is the server's constructor.
func New() *Server {
	return &Server{
		ph: handlers.NewPWD(),
	}
}

// Set prepares and sets the server to run.
func (s *Server) Set(cfg *config.Config) error {

	// create a new router
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// init the password handler
	if err := s.ph.Init(cfg.DBConfig); err != nil {
		return errors.Wrap(err, "failed to start password handler")
	}

	// apply routing
	SetRoutes(r, s.ph)

	// set the server properties
	s.srv = &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           r,
		ReadTimeout:       1000 * time.Millisecond,
		ReadHeaderTimeout: 500 * time.Millisecond,
		WriteTimeout:      1000 * time.Millisecond,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	return nil
}

// Start starts the server.
func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

// Stop closes all connections and dials.
func (s *Server) Stop() error {
	// stop the handler's services
	return s.ph.Close()
}
