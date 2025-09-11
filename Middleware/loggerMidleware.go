package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()
		
		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		fmt.Printf("[%s] %s %s %d %v\n",
			end.Format(time.RFC3339),
			c.Request.Method,
			c.Request.URL.Path,
			statusCode,
			latency,
		)
	}
}
