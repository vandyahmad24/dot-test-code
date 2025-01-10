package middleware

import (
	"dot-test-vandy/lib/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic recovered: %v\n", err)

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
				logger.MakeLogEntry(c).Error(fmt.Sprintf("Panic recovered: %v", err))
				c.Abort()
			}
		}()

		c.Next()
	}
}
