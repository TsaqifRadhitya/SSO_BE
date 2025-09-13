package Routes

import (
	"SSO_BE_API/Handler/Application"
	"SSO_BE_API/Middleware"
	"github.com/gin-gonic/gin"
)

func ApplicationRoutes(c *gin.RouterGroup) {
	clientRoutes := c.Group("/application")

	//Set auth middleware to application route
	clientRoutes.Use(Middleware.AuthMiddleware())

	//get all application
	clientRoutes.GET("/", Application.IndexApplicationHandler())

	//create new application
	clientRoutes.POST("/create", Application.StoreApplicationHandler())

	//get application configuration
	clientRoutes.GET(":id", Application.ShowApplicationHandler())

	//generate new client key for existed application
	clientRoutes.GET(":id/refresh", Application.RefreshApplicationKeyHandler())

	//delete existed application
	clientRoutes.DELETE("/:id", Application.DeleteApplicationHandler())

	//add new white list callback url on existed application
	clientRoutes.POST("/:id/callback", Application.StoreApplicationCallbackHandler())

	//update white list callback url on existed application
	clientRoutes.PATCH("/:id/:callback_id", Application.UpdateApplicationCallbackHandler())

	//delete white list callback url on existed application
	clientRoutes.DELETE("/:id/:callback_id", Application.DeleteApplicationCallbackHandler())

}
