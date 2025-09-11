package Middleware

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		BearerToken := c.GetHeader("Authorization")

		if BearerToken == "" {
			c.JSON(401, gin.H{
				"status":  http.StatusBadRequest,
				"message": "missing Authorization header",
			})
			c.Abort()
			return

		}

		parts := strings.SplitN(BearerToken, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "required credentials are: Bearer",
			})
			c.Abort()
			return
		}
		Credential, err := Utils.Claims(parts[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		if err := Config.DB.
			Where("jwt_token = ? AND jwt_expiry > ?", Credential, time.Now()).First(&Entity.Session{}).Error; err != nil {
			errMsg := err.Error()
			c.JSON(http.StatusUnauthorized, Response.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   &errMsg,
			})
			c.Abort()
			return
		}

		c.Set("User", Credential.UserCredential)
		c.Next()
	}
}
