package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronalc90/skillpath/internal/dto"
)

// Recovery handles panics and returns a 500 error response.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC RECOVERY] %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
					Error:   "internal_error",
					Message: "An unexpected error occurred",
				})
			}
		}()
		c.Next()
	}
}
