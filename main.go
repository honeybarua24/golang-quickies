package main

import (
	"my-gin-api/database"
	"my-gin-api/logger"
	"my-gin-api/middleware"
	"my-gin-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()
	r := gin.New()
	logger.Logger.Info("Starting server")

	//Initialize database connection
	database.InitDB()

	//Use middleware for various contexts
	r.Use(middleware.RequestIDMiddleware())

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Info("Error Starting Server")
	}
}
