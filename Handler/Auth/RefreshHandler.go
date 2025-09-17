package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	Auth2 "SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		refresh_token, err := c.Cookie("refresh_token")

		fmt.Println(refresh_token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, DTOResponse.ResponseError[interface{}]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		newCredential, err := Auth2.RefreshTokenService(refresh_token)

		if err != nil {
			formatedError := Utils.ErrorFormater(err)
			c.JSON(formatedError.Status, formatedError)
			c.Abort()
			return
		}

		c.SetCookie(
			"refresh_token",            // nama cookie
			newCredential.RefreshToken, // value
			60*60*24*30,                // maxAge 30 hari (detik)
			"/",                        // path
			"localhost",                // domain
			false,                      // Secure=false untuk dev (HTTPS nanti true)
			true,                       // HttpOnly
		)

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[DTOAuth.Auth]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    newCredential,
		})
	}
}
