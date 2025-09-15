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

		var Session Entity.Session

		if err := Config.DB.Where("jwt_token = ?", Jwt).First(&Session).Error; err != nil {
			c.JSON(http.StatusUnauthorized, DTOResponse.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		fmt.Println(Session)

		if time.Now().After(Session.JwtExpiry) {
			c.JSON(http.StatusUnauthorized, DTOResponse.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		c.Set("User", fmt.Sprintf("%v", Credential.UserCredential))
		c.Next()
	}
}
