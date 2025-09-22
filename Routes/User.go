package Routes

import (
	"SSO_BE_API/Handler/User"
	"SSO_BE_API/Middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.RouterGroup) {
	userRoutes := c.Group("/user")

	//get user information via verify token
	userRoutes.GET("/verify", User.GetUserByVerifyTokenHandler())

	//get user information via jwt token
	userRoutes.GET("/", Middleware.AuthMiddleware(), User.GetUserByJwtTokenHandler())

	userRoutes.GET("/access_log", Middleware.AuthMiddleware(), User.GetAccessLogHandler())

	userRoutes.GET("/connected_app", Middleware.AuthMiddleware(), User.GetConnectedAppHandler())
}
