package Routes

import (
	"SSO_BE_API/Handler/Auth"
	"SSO_BE_API/Middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(c *gin.RouterGroup) {
	authRoutes := c.Group("/auth")
	{
		//login sso account
		authRoutes.POST("/login", Auth.LoginHandler())

		//logout sso account
		authRoutes.POST("/logout", Middleware.AuthMiddleware(), Auth.LogoutHandler())

		//register new sso account
		authRoutes.POST("/register", Auth.RegisterHandler())

		//refresh jwt token
		authRoutes.POST("/refresh", Auth.RefreshTokenHandler())

		//verify Client Key and Callback URL before giving permision to use SSO
		authRoutes.POST("/verify_access", Auth.VerifyAccessHandler())

		//sign in to consumer application
		authRoutes.POST("/sso", Middleware.AuthMiddleware(), Auth.SSOHandler())
	}
}
