package User

import (
	DTO "SSO_BE_API/Model/DTO/Response"
	DTOUser "SSO_BE_API/Model/DTO/User"
	"SSO_BE_API/Service/User"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResetPasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resetPasswordRequest DTOUser.ResetPassword

		if err := c.ShouldBind(&resetPasswordRequest); err != nil {
			c.JSON(400, DTO.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(resetPasswordRequest); err != nil {
			c.JSON(400, DTO.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		if err := User.ResetPasswordService(resetPasswordRequest); err != nil {
			formatedError := Utils.ErrorFormater(err)
			c.JSON(formatedError.Status, formatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, DTO.ResponseSuccess[interface{}]{
			Status:  http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		})
	}
}
