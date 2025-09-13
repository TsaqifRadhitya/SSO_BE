package Middleware

import (
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case error:
					c.JSON(http.StatusInternalServerError, DTOResponse.ResponseError[string]{
						Status:  http.StatusInternalServerError,
						Message: http.StatusText(http.StatusInternalServerError),
						Error:   err.(error).Error(),
					})
					c.Abort()
					return
				default:
					return
				}
			}
		}()
		c.Next()
	}
}
