package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization basic token required"})
			c.Abort()
			return
		}

		const (
			expectedUsername = "admin"
			expectedPassword = "#Admin123"
		)

		isValid := (username == expectedUsername) && (password == expectedPassword)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
