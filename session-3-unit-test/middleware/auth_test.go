package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"training-golang/session-3-unit-test/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_PositiveCase(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	req, _ := http.NewRequest("GET", "/private", nil)
	req.Header.Set("Authorization", "valid-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Private data"}`, w.Body.String())
}

func TestAuthMiddleware_NegativeCase_NoToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	req, _ := http.NewRequest("GET", "/private", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization token required")
}

func TestAuthMiddleware_NegativeCase_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	req, _ := http.NewRequest("GET", "/private", nil)
	req.Header.Set("Authorization", "invalid-token")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid authorization token")
}
