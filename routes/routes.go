package routes

import (
	"first_project/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	// Create v1 route group
	v1 := router.Group("/api/v1")
	{
		// Register the HomeController route within the v1 group
		v1.GET("/home", controllers.UserController)
		v1.GET("/home/:id", controllers.UserController)
		v1.POST("/home", controllers.UserController)
		v1.DELETE("/home/:id", controllers.UserController)

		// You can add more routes within the v1 group here
		// Example: v1.POST("/users", controllers.CreateUser)
	}

	return router
}
