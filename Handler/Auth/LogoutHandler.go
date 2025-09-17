package Auth

import (
	"SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		refresh_token, err := c.Cookie("refresh_token")

		fmt.Println(refresh_token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, DTO.ResponseError[interface{}]{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		if err := Auth.LogoutService(refresh_token); err != nil {
			formatedErr := Utils.ErrorFormater(err)
			c.JSON(formatedErr.Status, formatedErr)
			c.Abort()
			return
		}

		c.SetCookie(
			"refresh_token", // nama cookie
			"",              // kosongin value
			-1,              // maxAge negatif biar browser hapus cookie
			"/",             // path
			"localhost",     // domain (samain kayak waktu set)
			false,           // secure (true kalau HTTPS)
			true,            // httpOnly
		)

		c.JSON(http.StatusOK, DTO.ResponseSuccess[*interface{}]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		})
	}
}
