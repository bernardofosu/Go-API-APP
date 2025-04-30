// package userRoutes

// import (
// 	"github.com/gin-gonic/gin"
// 	"my-go-gin-app/controllers"
// )

// func RegisterUserRoutes(r *gin.Engine) {
// 	userGroup := r.Group("/api/users")
// 	{
// 		userGroup.GET("/", controllers.GetUsers)
// 	}
// }

package userRoutes

import (
	"github.com/gin-gonic/gin"
	"my-go-gin-app/controllers"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
	}
}
