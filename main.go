package main

import (
	"SSO_BE_API/Config"
)

func main() {
	if err := Config.DbConnect(); err != nil {
		panic(err.Error())
	}
	defer Config.DbClose()

	server := GetServer()

	if err := server.Run(); err != nil {
		panic(err.Error())
	}
}
