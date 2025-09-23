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

	userRoutes.POST("/reset_password", User.ResetPasswordHandler())

	TwoVerificationRouteGroup := userRoutes.Group("/TwoVerification", Middleware.AuthMiddleware())
	{
		TwoVerificationRouteGroup.GET("/", User.GetOTPHandler())
		TwoVerificationRouteGroup.POST("/", User.Activate2FAHandler())
		TwoVerificationRouteGroup.DELETE("/", User.Remove2FAHandler())
	}
}
