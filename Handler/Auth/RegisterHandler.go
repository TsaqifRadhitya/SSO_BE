package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerPayload DTOAuth.Register

		if err := c.ShouldBind(&registerPayload); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseSuccess[string]{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(registerPayload); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		ress, err := Auth.RegisterService(registerPayload)

		if err != nil {
			formatedError := Utils.ErrorFormater(err)
			c.JSON(formatedError.Status, formatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[Entity.User]{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    ress,
		})

	}
}
