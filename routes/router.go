package routes

import (
	"github.com/gin-gonic/gin"
	"my-go-gin-app/controllers"
	"my-go-gin-app/routes/userRoutes"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// Load templates and static files
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	// Basic HTML rendering
	router.GET("/", controllers.RenderUsersPage)

	// Base API group
	api := router.Group("/api")

	userRoutes.RegisterUserRoutes(api) // 🔗 All /api/users/* go through here

	return router
}
