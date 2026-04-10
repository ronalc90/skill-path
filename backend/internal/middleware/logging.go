package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logging logs each HTTP request with method, path, status, and duration.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		if query != "" {
			path = path + "?" + query
		}

		log.Printf("[%s] %d | %v | %s", c.Request.Method, status, duration, path)
	}
}
