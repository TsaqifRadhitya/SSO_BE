package Handler

import (
	"SSO_BE_API/Model"
	"SSO_BE_API/Service"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var GetUserRequest Model.VerifyToken
		if err := c.BindJSON(&GetUserRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := Utils.Validate(&GetUserRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		data, err := Service.GetUserService(GetUserRequest)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": http.StatusText(http.StatusOK),
			"data":    data,
		})

	}
}
