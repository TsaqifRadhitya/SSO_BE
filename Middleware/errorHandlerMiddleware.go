package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case error:
					c.JSON(500, gin.H{
						"status":  500,
						"error":   err.(error).Error(),
						"message": http.StatusText(500),
					})
				}
			}
			return
		}()
		c.Next()
	}
}
