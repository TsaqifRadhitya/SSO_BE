package Application

import (
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Service/Application"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshApplicationKeyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApplicationOwnerCredential, _ := c.Get("user")
		ApplicationId := c.Query("id")

		RefreshApplicationKeyRequest := DTOApplication.DeleteApplicationCallaback{
			ApplicationId: ApplicationId,
			OwnerId:       ApplicationOwnerCredential.(string),
		}

		if err := Utils.Validate(RefreshApplicationKeyRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		if err := Application.DeleteApplicationCallbackService(RefreshApplicationKeyRequest); err != nil {
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
