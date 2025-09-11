package Handler

import (
	"SSO_BE_API/Model"
	"SSO_BE_API/Service"
	"SSO_BE_API/Utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var LoginRequest Model.LoginRequest
		if err := c.ShouldBind(&LoginRequest); err != nil {
			panic(err)
		}

		if validateRequest := Utils.Validate(LoginRequest); validateRequest != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "Validation Error",
				"error":   validateRequest,
			})
			return
		}

		CallbackUrl, err := Service.LoginService(LoginRequest)
		if err != nil {
			c.JSON(404, gin.H{})
			return
		}

		c.JSON(200, gin.H{
			"status":  200,
			"message": "success",
			"data": gin.H{
				"callback_url": CallbackUrl,
			},
		})
	}
}
