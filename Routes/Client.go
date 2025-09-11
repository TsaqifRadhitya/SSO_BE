package Routes

import "github.com/gin-gonic/gin"

func ApplicationRoutes(c *gin.RouterGroup) {
	clientRoutes := c.Group("/application")

	//get all application
	clientRoutes.GET("/")

	//get application configuration
	clientRoutes.GET(":id")

	//generate new client key for existed application
	clientRoutes.GET(":id/refresh")

	//create new application
	clientRoutes.POST("/create")

	//delete existed application
	clientRoutes.DELETE("/:id")

	//add new white list callback url on existed application
	clientRoutes.POST("/:id/callback")

	//update white list callback url on existed application
	clientRoutes.PATCH("/:id/:callback_id")

	//delete white list callback url on existed application
	clientRoutes.DELETE("/:id/:callback_id")

}
