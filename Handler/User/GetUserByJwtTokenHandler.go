package User

import (
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Service/User"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByJwtTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := c.Get("User")

		data, err := User.GetUserByJwtTokenService(userId.(string))

		if err != nil {
			formaterError := Utils.ErrorFormater(err)
			c.JSON(formaterError.Status, formaterError)
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
