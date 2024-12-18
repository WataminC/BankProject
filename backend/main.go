package main

import (
	"bank/config"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Print("Configuration successful\n")
}
