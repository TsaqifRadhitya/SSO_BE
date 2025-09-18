package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	DTOReponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginPayload DTOAuth.Login

		if err := c.ShouldBind(&loginPayload); err != nil {
			c.JSON(http.StatusBadRequest, DTOReponse.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(loginPayload); err != nil {
			c.JSON(http.StatusBadRequest, DTOReponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		result, err := Auth.LoginService(loginPayload)

		if err != nil {
			formatedErr := Utils.ErrorFormater(err)
			c.JSON(formatedErr.Status, formatedErr)
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
			"refresh_token",     // nama cookie
			result.RefreshToken, // value
			60*60*24*30,         // maxAge 30 hari (detik)
			"/",                 // path
			domain,              // domain
			true,                // Secure=false untuk dev (HTTPS nanti true)
			true,                // HttpOnly
		)

		c.JSON(http.StatusOK, DTOReponse.ResponseSuccess[DTOAuth.Auth]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    result,
		})
	}
}
