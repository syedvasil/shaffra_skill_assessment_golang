package middleware_test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/syedvasil/shaffra_skill_assessment_golang/internal/middleware"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLogProcessTime(t *testing.T) {
	// Create a buffer to capture log output
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf) // Set log output to buffer
	defer func() {
		log.SetOutput(nil) // Reset log output after the test
	}()

	w := CreateRequestAndCall()

	// status Checks
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())

	// Checks to verify the logs
	logOutput := logBuf.String()
	assert.True(t, strings.Contains(logOutput, "Request processed in"), "Log output does not contain expected message")
	assert.True(t, strings.Contains(logOutput, "ms") || strings.Contains(logOutput, "µs"), "Log output does not contain a valid duration")
}

func CreateRequestAndCall() *httptest.ResponseRecorder {
	// Create a new Gin router
	router := gin.New()
	router.Use(middleware.LogProcessTime()) // Use the middleware

	// temp route that with "OK" status
	router.GET("/test", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond) // Simulate processing time
		c.String(http.StatusOK, "OK")
	})

	// Create a test request and perform
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)
	return w
}
