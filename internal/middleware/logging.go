package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LogProcessTime() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Start time
		start := time.Now()

		c.Next()

		// Calculate time taken
		duration := time.Since(start)
		log.Printf("Request processed in %v\n", duration)
	}
}
