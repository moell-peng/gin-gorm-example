package main

import (
	"gin-gorm-example/routes"
	"gin-gorm-example/database"
	"gin-gorm-example/config"
	"fmt"
)

func main() {
	if err := config.Load("config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	_, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}

	router := routes.InitRouter()
	router.Run(config.Get().Addr)
}

