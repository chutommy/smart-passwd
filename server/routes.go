package server

import (
	handlers "github.com/chutified/smart-passwd/handlers"
	"github.com/gin-gonic/gin"
)

// SetRoutes sets all routing for the gin's engine.
func SetRoutes(e *gin.Engine, ph *handlers.PWDhandler) {

	// ping to test server status
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"response": "pong"})
	})

	// set API
	api := e.Group("/api")
	{
		api.POST("/passwd", ph.PasswordGen)
	}

	// main page
	e.Static("/assets", "./templates/assets")
	e.LoadHTMLFiles("templates/index.html")
	e.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
}
