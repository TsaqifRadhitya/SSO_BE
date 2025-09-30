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

func VerifyResetPasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var verifyResetPasswordRequest DTOUser.VerifyResetPassword

		reseted_mail, err := c.Cookie("reseted_mail")

		otp, err := c.Cookie("otp")

		if err != nil {
			c.JSON(http.StatusUnauthorized, DTO.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		if err := c.ShouldBind(&verifyResetPasswordRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTO.ResponseError[*interface{}]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(verifyResetPasswordRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTO.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		if isValid := User.VerifyResetPasswordService(verifyResetPasswordRequest, DTOUser.VerifyResetPassword{otp, reseted_mail}); !isValid {
			c.JSON(http.StatusUnauthorized, DTO.ResponseError[string]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
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
			"otp",  // nama cookie
			"",     // kosongin value
			-1,     // maxAge negatif biar browser hapus cookie
			"/",    // path
			domain, // domain (samain kayak waktu set)
			true,   // secure (true kalau HTTPS)
			true,   // httpOnly
		)

		c.SetCookie(
			"reseted_mail", // nama cookie
			"",             // kosongin value
			-1,             // maxAge negatif biar browser hapus cookie
			"/",            // path
			domain,         // domain (samain kayak waktu set)
			true,           // secure (true kalau HTTPS)
			true,           // httpOnly
		)

		c.SetCookie(
			"reset_password_email_access", // nama cookie
			reseted_mail,                  // value
			60*15,                         // maxAge 30 hari (detik)
			"/",                           // path
			domain,                        // domain
			true,                          // Secure=false untuk dev (HTTPS nanti true)
			true,                          // HttpOnly
		)

		c.JSON(http.StatusOK, DTO.ResponseSuccess[string]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		})
	}
}
