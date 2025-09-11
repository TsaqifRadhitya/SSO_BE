package Routes

import "github.com/gin-gonic/gin"

func UserRoutes(c *gin.RouterGroup) {
	userRoutes := c.Group("/user")

	//get user information with access token
	userRoutes.GET("/")
}
