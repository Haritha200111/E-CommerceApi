package routes

import (
	"E-COMMERCEAPI/controllers"
	"E-COMMERCEAPI/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// Public route for login (authentication)
	api.POST("/login", controllers.Login)

	// Secured routes (require JWT authentication)
	api.Use(middleware.JWTAuthMiddleware())
	{
		// Category routes
		api.GET("/categories", controllers.GetCategories)         // Get all categories
		api.POST("/categories", controllers.CreateCategory)       // Create a new category
		api.GET("/categories/:id", controllers.GetCategoryByID)   // Get a category by ID
		api.PUT("/categories/:id", controllers.UpdateCategory)    // Update a category by ID
		api.DELETE("/categories/:id", controllers.DeleteCategory) // Delete a category by ID

		// Product routes
		api.GET("/products", controllers.GetProducts)          // Get all products
		api.POST("/products", controllers.CreateProduct)       // Create a new product
		api.GET("/products/:id", controllers.GetProductByID)   // Get a product by ID
		api.PUT("/products/:id", controllers.UpdateProduct)    // Update a product by ID
		api.DELETE("/products/:id", controllers.DeleteProduct) // Delete a product by ID

		// Variant routes
		api.GET("/variants", controllers.GetVariants)          // Get all variants
		api.POST("/variants", controllers.CreateVariant)       // Create a new variant
		api.GET("/variants/:id", controllers.GetVariantByID)   // Get a variant by ID
		api.PUT("/variants/:id", controllers.UpdateVariant)    // Update a variant by ID
		api.DELETE("/variants/:id", controllers.DeleteVariant) // Delete a variant by ID

		// Order routes
		api.POST("/orders", controllers.CreateOrder)     // Create a new order
		api.GET("/orders/:id", controllers.GetOrderByID) // Get an order by ID
		api.GET("/orders", controllers.GetOrders)        // Get all orders
	}
}
