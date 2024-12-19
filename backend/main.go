package main

import (
	"bank/config"
	"bank/router"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Print("Configuration successful\n")

	port := config.AppConfig.App.Port

	if port == "" {
		port = "8080"
	}

	r := router.SetRouter()
	r.Run(":" + port)
}
