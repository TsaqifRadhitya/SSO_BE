package User

import (
	DTOResponse "SSO_BE_API/Model/DTO/Response"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Service/User"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAccessLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := c.Get("User")

		data, err := User.GetAccesLoggerService(userId.(string))

		if err != nil {
			formatedError := Utils.ErrorFormater(err)
			c.JSON(formatedError.Status, formatedError)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, DTOResponse.ResponseSuccess[[]Entity.AccessLog]{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    data,
		})
	}
}
