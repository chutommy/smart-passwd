package server

import (
	"errors"
	"net/http"

	"github.com/chutified/smart-passwd/pkg/engine"
	"github.com/gin-gonic/gin"
)

// homePageHandler serves the home page.
func homePageHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.HTML(200, "index.html", nil)
	}
}

// pingHandler serves ping requests.
func pingHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "online"})
	}
}

// passwordGenHandler returns a gin.HandlerFunc of the password
// generation handler.
func passwordGenHandler(e *engine.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		resp, err := e.Generate(c, engine.NewRequest(
			int16(req.Length),
			int16(req.ExtraSecurity),
			req.Helper,
		))
		if err != nil {
			code := http.StatusInternalServerError
			if errors.Is(err, engine.ErrInvalidRequirements) {
				code = http.StatusBadRequest
			}

			c.JSON(code, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, GenResponse{
			Passwd: resp.Password(),
			Helper: resp.Helper(),
		})
	}
}
