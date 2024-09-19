package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"training-golang/session-3-unit-test/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHelloMessage(t *testing.T) {
	t.Run("Positive case - correct message", func(t *testing.T) {
		expectOutput := "Hello from gin"
		actualOutput := handler.GetHelloMessage()

		require.Equal(t, expectOutput, actualOutput, "The message should be '%s'", expectOutput)
	})
}

func TestRootHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/", handler.RootHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"message":"Hello from gin"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

type JsonRequest struct {
	Message string `"json:message"`
}

func TestPostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/", handler.PostHandler)

	t.Run("Positive case", func(t *testing.T) {
		requestBody := JsonRequest{Message: "Hello from test"}
		requestBodyBytes, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		expectedBody := `{"message":"Hello from test"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("Negative case - EOF Error", func(t *testing.T) {
		requestBody := ""
		requestBodyBytes := []byte(requestBody)

		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.Contains(t, w.Body.String(), "{\"error\":\"EOF\"}")
	})
}
