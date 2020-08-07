package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

	// apply crach free middleware
	r.Use(gin.Recovery())

	// apply custom logging
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\" %s %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// log to file and os.Stdout
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

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
