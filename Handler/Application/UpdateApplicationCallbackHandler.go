package Application

import (
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Service/Application"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateApplicationCallbackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApplicationOwnerCredential, _ := c.Get("User")
		ApplicationId := c.Param("id")
		CallbackId := c.Param("callback_id")

		var UpdateApplicationCallbackRequest DTOApplication.UpdateApplicationCallaback

		if err := c.ShouldBind(&UpdateApplicationCallbackRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		UpdateApplicationCallbackRequest.OwnerId = ApplicationOwnerCredential.(string)
		UpdateApplicationCallbackRequest.ApplicationId = ApplicationId
		UpdateApplicationCallbackRequest.CallbackId = CallbackId

		if err := Utils.Validate(UpdateApplicationCallbackRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}
		data, err := Application.UpdateApplicationCallbackService(UpdateApplicationCallbackRequest)
		if err != nil {
			FormatedError := Utils.ErrorFormater(err)
			c.JSON(FormatedError.Status, FormatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[Entity.CallbackApplication]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    data,
		})
	}
}
