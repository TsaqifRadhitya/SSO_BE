package Application

import (
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Service/Application"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApplicationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApplicationOwnerCredential, _ := c.Get("user")

		data, err := Application.GetAllApplicationService(ApplicationOwnerCredential.(string))

		if err != nil {
			FormatedError := Utils.ErrorFormater(err)
			c.JSON(FormatedError.Status, FormatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[[]Entity.Application]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    data,
		})
	}
}
