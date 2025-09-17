package main

import (
	"SSO_BE_API/Config"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	if err := Config.DbConnect(); err != nil {
		panic(err.Error())
	}
	defer Config.DbClose()

	server := GetServer()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // asal FE kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	if err := server.Run(); err != nil {
		panic(err.Error())
	}
}
