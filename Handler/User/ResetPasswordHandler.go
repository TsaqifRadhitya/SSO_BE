package User

import (
	DTO "SSO_BE_API/Model/DTO/Response"
	DTOUser "SSO_BE_API/Model/DTO/User"
	"SSO_BE_API/Service/User"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func ResetPasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resetPasswordRequest DTOUser.ResetPassword

		if err := c.ShouldBind(&resetPasswordRequest); err != nil {
			c.JSON(400, DTO.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(resetPasswordRequest); err != nil {
			c.JSON(400, DTO.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}
		otp, err := User.ResetPasswordService(resetPasswordRequest)
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
			"otp",             // nama cookie
			strconv.Itoa(otp), // value
			60*5,              // maxAge 30 hari (detik)
			"/",               // path
			domain,            // domain
			true,              // Secure=false untuk dev (HTTPS nanti true)
			true,              // HttpOnly
		)

		c.SetCookie(
			"reseted_mail",             // nama cookie
			resetPasswordRequest.Email, // value
			60*15,                      // maxAge 30 hari (detik)
			"/",                        // path
			domain,                     // domain
			true,                       // Secure=false untuk dev (HTTPS nanti true)
			true,                       // HttpOnly
		)

		c.JSON(http.StatusCreated, DTO.ResponseSuccess[interface{}]{
			Status:  http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		})
	}
}
