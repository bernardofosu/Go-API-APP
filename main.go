package main

import (
	"fmt"
	"log"
	"my-go-gin-app/config"
	"my-go-gin-app/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	router := routes.InitRoutes()

	port := config.GetPort()
	fmt.Println("🚀 Starting server on port:", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}

	for _, ri := range router.Routes() {
		fmt.Printf("🔗 Registered route: %s %s\n", ri.Method, ri.Path)
	}

}
