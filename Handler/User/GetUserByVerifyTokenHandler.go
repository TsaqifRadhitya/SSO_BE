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

func GetUserByVerifyTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		verifyToken := c.Query("verify_token")

		if verifyToken == "" {
			c.JSON(400, DTOResponse.ResponseError[string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   "required verify token",
			})
			c.Abort()
			return
		}

		verifyTokenDTO := DTOUser.VerifyToken{
			Token: verifyToken,
		}

		if err := Utils.Validate(verifyTokenDTO); err != nil {
			c.JSON(http.StatusBadRequest, DTOResponse.ResponseError[map[string]string]{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err,
			})
			c.Abort()
			return
		}

		data, err := User.GetUserByVerifyTokenService(verifyTokenDTO)

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
