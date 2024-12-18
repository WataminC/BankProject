package main

import (
	"bank/config"
	"bank/router"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Print("Configuration successful\n")

	r := router.SetRouter()
	r.Run()
}
