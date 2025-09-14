package main

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Handler"
	"SSO_BE_API/Middleware"
	"SSO_BE_API/Routes"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	if Config.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	//server.Use(Middleware.LoggerMiddleware())
	server.Use(Middleware.ErrorMiddleware())

	//set routing group prefix /api
	api := server.Group("/api")
	{
		Routes.AuthRoutes(api)
		Routes.UserRoutes(api)
		Routes.ApplicationRoutes(api)
	}

	server.NoRoute(Handler.NotFoundRouteHandler())

	return server
}
