package Auth

import (
	Auth2 "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Service/Auth"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerPayload Auth2.RegisterJson

		if err := c.ShouldBindJSON(&registerPayload); err != nil {
			c.JSON(http.StatusBadRequest, Response.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(registerPayload); err != nil {
			c.JSON(http.StatusBadRequest, Response.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   &err,
			})
			c.Abort()
			return
		}

		ress, err := Auth.RegisterService(registerPayload)

		if err != nil {
			msg := err.Error()
			c.JSON(http.StatusInternalServerError, Response.ResponseError[string]{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   &msg,
			})
		}

		c.JSON(http.StatusOK, Response.ResponseSuccess[Entity.User]{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    &ress,
		})

	}
}
