package User

import (
	DTO "SSO_BE_API/Model/DTO/Response"
	DTOUser "SSO_BE_API/Model/DTO/User"
	"SSO_BE_API/Service/User"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetNewPasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, err := c.Cookie("reset_password_email_access")

		if err != nil {
			c.JSON(http.StatusUnauthorized, DTO.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		SetNewPasswordRequest := DTOUser.SetNewPassword{
			Email: email,
		}

		if err := c.ShouldBind(&SetNewPasswordRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTO.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(SetNewPasswordRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTO.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		if err := User.SetNewPasswordService(SetNewPasswordRequest); err != nil {
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
			"reset_password_email_access", // nama cookie
			"",                            // kosongin value
			-1,                            // maxAge negatif biar browser hapus cookie
			"/",                           // path
			domain,                        // domain (samain kayak waktu set)
			true,                          // secure (true kalau HTTPS)
			true,                          // httpOnly
		)

		c.JSON(http.StatusOK, DTO.ResponseSuccess[string]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		})
	}
}
