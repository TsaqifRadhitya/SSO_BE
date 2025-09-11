package Routes

import (
	"SSO_BE_API/Handler/User"
	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.RouterGroup) {
	userRoutes := c.Group("/user")

	//get user information with access token
	userRoutes.GET("/", User.GetUserHandler())
}
