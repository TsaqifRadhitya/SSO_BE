package User

import (
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	DTOUser "SSO_BE_API/Model/DTO/User"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Service/User"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getUserRequest DTOUser.VerifyToken

		if err := c.ShouldBindJSON(&getUserRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		if err := Utils.Validate(getUserRequest); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		data, err := User.GetUserService(getUserRequest)

		if err != nil {
			FormatedError := Utils.ErrorFormater(err)
			c.JSON(FormatedError.Status, FormatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[Entity.User]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    data,
		})
	}
}
