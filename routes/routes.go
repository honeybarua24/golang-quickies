package routes

import (
	"my-gin-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Group routes for versioning
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", handlers.GetUsers)
		v1.GET("/users/:id", handlers.GetUserByID)
		v1.POST("/users", handlers.CreateUser)
		v1.PUT("/users/:id", handlers.UpdateUser)
		v1.DELETE("/users/:id", handlers.DeleteUser)
	}
}
