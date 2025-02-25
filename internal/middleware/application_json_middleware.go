package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ApplicationJsonResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.Path, "/swagger/") {
			c.Writer.Header().Set("Content-Type", "application/json")
		}

		c.Next()
	}
}
