package main

import (
	"SSO_BE_API/Handler"
	"SSO_BE_API/Middleware"
	"SSO_BE_API/Routes"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	server := gin.Default()
	server.Use(Middleware.Logger())
	server.Use(Middleware.ErrorHandler())
	api := server.Group("/api")
	{
		api.POST("/login", Handler.LoginHandler())
		api.GET("/user", Handler.GetUserHandler())
		Routes.AuthRoutes(api)
		Routes.UserRoutes(api)
		Routes.ApplicationRoutes(api)
	}

	return server
}
