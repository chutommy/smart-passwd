package server

import (
	handlers "github.com/chutified/smart-passwd/handlers"
	"github.com/gin-gonic/gin"
)

// SetRoutes sets all routing for the gin's engine.
func SetRoutes(e *gin.Engine, ph *handlers.PWDhandler) {

	// ping to test server status
	e.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// set API
	api := e.Group("/api")
	{
		api.POST("/passwd", ph.PasswordGen)
	}
}
