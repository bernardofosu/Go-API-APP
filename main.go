// package main

// import (
// 	"log"
// 	"my-go-gin-app/config"
// 	"my-go-gin-app/routes"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	config.LoadEnv()
// 	router := gin.Default()

// 	routes.RegisterRoutes(router)

// 	port := config.AppConfig.Port
// 	log.Printf("Server running on port %s", port)
// 	router.Run(":" + port)
// }

// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"my-go-gin-app/config"
// 	"my-go-gin-app/routes"
// )

// func main() {
// 	config.LoadEnv()
// 	config.ConnectDB()

// 	router := gin.Default()
// 	routes.RegisterRoutes(router)

// 	router.Run(":" + config.GetPort())
// }

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
}
