package routes

import (
	"ecommerce/controllers"
	"ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// Public route for login (authentication)
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)

	// Public routes for Category
	api.GET("/categories", controllers.GetCategories)       // Get all categories
	api.POST("/category", controllers.CreateCategory)       // Create a new category
	api.POST("/subcategory", controllers.CreateSubCategory) // Create a new category
	api.GET("/category", controllers.GetCategoryByID)       // Get a category by ID
	api.PUT("/categories", controllers.UpdateCategory)      // Update a category by ID
	api.DELETE("/category", controllers.DeleteCategory)     // Delete a category by ID

	// Public routes for Product
	api.GET("/products", controllers.GetProducts)     // Get all products
	api.POST("/products", controllers.CreateProduct)  // Create a new product
	api.GET("/product", controllers.GetProductByID)   // Get a product by ID
	api.PUT("/product", controllers.UpdateProduct)    // Update a product by ID
	api.DELETE("/product", controllers.DeleteProduct) // Delete a product by ID

	// Public routes for Variant
	api.GET("/variants", controllers.GetVariants)     // Get all variants
	api.POST("/variant", controllers.CreateVariant)   // Create a new variant
	api.GET("/variant", controllers.GetVariantById)   // Get a variant by ID
	api.PUT("/variant", controllers.UpdateVariant)    // Update a variant by ID
	api.DELETE("/variant", controllers.DeleteVariant) // Delete a variant by ID

	// Order routes (JWT protected)
	orderRoutes := api.Group("/orders")
	orderRoutes.Use(middleware.JWTAuthMiddleware())
	{
		orderRoutes.POST("/create", controllers.CreateOrder) // Create a new order
		// orderRoutes.GET("/:id", controllers.GetOrderByID) // Get an order by ID
		orderRoutes.GET("/getorders", controllers.GetOrders) // Get all orders
	}
}
