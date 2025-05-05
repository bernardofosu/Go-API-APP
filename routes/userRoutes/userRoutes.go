package userRoutes

import (
	"github.com/gin-gonic/gin"
	"my-go-gin-app/controllers"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users") // Matches /api/users/*
	{
		userGroup.GET("/", controllers.GetUsers)         // GET /api/users/
		userGroup.POST("/", controllers.CreateUser)      // POST /api/users/
		userGroup.GET("/:id", controllers.GetUserByID)   // GET /api/users/:id
		userGroup.DELETE("/:id", controllers.DeleteUser) // DELETE /api/users/:id
		userGroup.PUT("/:id", controllers.UpdateUser)    // PUT /api/users/:id
		userGroup.PATCH("/:id", controllers.PatchUser)   // PATCH /api/users/:id

	}
}

// /api                  <- top-level group from router.go
//    /users             <- sub-group from userRoutes.go
//       GET /           <- route handled by GetUsers
