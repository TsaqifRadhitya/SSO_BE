package main

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Handler"
	"SSO_BE_API/Middleware"
	"SSO_BE_API/Routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func GetServer() *gin.Engine {
	if Config.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // domain Next.js
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true, // wajib supaya cookie dikirim
		MaxAge:           12 * time.Hour,
	}))
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
