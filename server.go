package main

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Middleware"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		var data []Entity.User
		Config.DB.Find(&data)

		c.JSON(http.StatusOK, gin.H{"data": data})

	})
	//server.Use(Middleware.LoggerMiddleware())
	server.Use(Middleware.ErrorMiddleware())
	api := server.Group("/api")
	{
		Routes.AuthRoutes(api)
		Routes.UserRoutes(api)
		Routes.ApplicationRoutes(api)
	}

	return server
}
