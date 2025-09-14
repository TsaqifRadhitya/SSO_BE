package Middleware

import (
	"SSO_BE_API/Config"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		BearerToken := c.GetHeader("Authorization")

		Jwt, err := Utils.ExtractBearerToken(BearerToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, DTOResponse.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		Credential, err := Utils.Claims(Jwt)

		if err != nil {
			c.JSON(http.StatusUnauthorized, DTOResponse.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		if err := Config.DB.Where("jwt_token = ? AND jwt_expiry > ?", Jwt, time.Now()).First(&Entity.Session{}).Error; err != nil {
			c.JSON(http.StatusUnauthorized, DTOResponse.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("User", fmt.Sprintf("%v", Credential.UserCredential))
		c.Next()
	}
}
