package Application

import (
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Application"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteApplicationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApplicationOwnerCredential, _ := c.Get("user")
		ApplicationId := c.Query("id")

		DeleteApplicationRequest := DTOApplication.DeleteApplication{
			ApplicationId: ApplicationId,
			OwnerId:       ApplicationOwnerCredential.(string),
		}

		if err := Utils.Validate(DeleteApplicationRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		if err := Application.DeleteApplicationService(DeleteApplicationRequest); err != nil {
			FormatedError := Utils.ErrorFormater(err)
			c.JSON(FormatedError.Status, FormatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[*interface{}]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		})
	}
}
