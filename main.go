package main

import (
	"ecommerce/config"
	// _ "ecommerce/docs" // Import the generated Swagger docs
	"ecommerce/routes"

	"github.com/gin-gonic/gin"
	// Swagger UI
)

// @title Ecommerce API
// @version 1.0
// @description This is the API documentation for the Ecommerce platform.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host localhost:8080
// @BasePath /api
func main() {
	// Initialize database connection
	config.ConnectDB()

	// Set up Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
