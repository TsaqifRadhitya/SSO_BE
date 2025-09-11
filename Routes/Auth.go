package Routes

import "github.com/gin-gonic/gin"

func AuthRoutes(c *gin.RouterGroup) {
	authRoutes := c.Group("/auth")
	{

		//login sso account
		authRoutes.POST("/login")

		//logout sso account
		authRoutes.GET("/logout")

		//register new sso account
		authRoutes.POST("/register")

		//refresh jwt token
		authRoutes.POST("/refresh")

		authRoutes.POST("/sso")
	}
}
