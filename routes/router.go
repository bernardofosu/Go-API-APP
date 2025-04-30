// package routes

// import (
// 	"github.com/gin-gonic/gin"
// 	"my-go-gin-app/routes/userRoutes"
// )

// func RegisterRoutes(r *gin.Engine) {
// 	userRoutes.RegisterUserRoutes(r)
// }

package routes

import (
	"github.com/gin-gonic/gin"
	"my-go-gin-app/routes/userRoutes"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	userRoutes.RegisterUserRoutes(api)

	return router
}
