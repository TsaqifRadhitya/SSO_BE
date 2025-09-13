package Auth

import (
	"SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		BearerToken := c.GetHeader("Authorization")
		parts := strings.SplitN(BearerToken, " ", 2)
		Auth.LogoutService(parts[1])
		c.JSON(http.StatusOK, DTO.ResponseSuccess[*interface{}]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		})
	}
}
