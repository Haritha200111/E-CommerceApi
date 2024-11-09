package main

import (
	"ecommerce-api/config"
	"ecommerce-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start server
	router.Run(":8080")
}
