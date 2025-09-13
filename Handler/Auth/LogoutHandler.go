package Auth

import (
	"SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		BearerToken := c.GetHeader("Authorization")
		parts := strings.SplitN(BearerToken, " ", 2)
		if err := Auth.LogoutService(parts[1]); err != nil {
			formatedErr := Utils.ErrorFormater(err)
			c.JSON(formatedErr.Status, formatedErr)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTO.ResponseSuccess[*interface{}]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		})
	}
}
