package main

import (
	"bank/config"
	"bank/global"
	"bank/router"
	"fmt"
)

func main() {
	config.InitConfig("./config")
	fmt.Print("Configuration successful\n")

	port := config.AppConfig.App.Port

	if port == "" {
		port = "8080"
	}

	defer global.RDB.Close()

	r := router.SetRouter()
	r.Run(":" + port)
}
