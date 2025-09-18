package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		method := c.Request.Method

		fmt.Printf("[CORS] %s %s from origin: %s\n", method, c.Request.URL.Path, origin)

		// Handle preflight OPTIONS request FIRST
		if method == "OPTIONS" {
			fmt.Println("[CORS] Handling OPTIONS preflight request")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT, PATCH, OPTIONS")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Continue to next middleware/handler
		c.Next()

		// Set CORS headers AFTER handler executed (untuk override apapun yang di-set handler)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT, PATCH, OPTIONS")
	})
}
