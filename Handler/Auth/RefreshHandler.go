package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	Auth2 "SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RefreshTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		refresh_token, err := c.Cookie("refresh_token")

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

		host := c.Request.Host

		var domain string
		if strings.Contains(host, "ngrok.io") {
			domain = ".ngrok.io" // wildcard untuk semua ngrok
		} else if strings.Contains(host, "localhost") {
			domain = "" // kosong untuk localhost
		} else {
			domain = "" // default
		}

		c.SetSameSite(http.SameSiteNoneMode)

		c.SetCookie(
			"refresh_token",            // nama cookie
			newCredential.RefreshToken, // value
			60*60*24*30,                // maxAge 30 hari (detik)
			"/",                        // path
			domain,                     // domain
			true,                       // Secure=false untuk dev (HTTPS nanti true)
			true,                       // HttpOnly
		)

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[DTOAuth.Auth]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    newCredential,
		})
	}
}
