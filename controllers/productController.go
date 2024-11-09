package controllers

import (
	"E-COMMERCEAPI/config"
	"E-COMMERCEAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts retrieves all products with their variants
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Preload("Variants").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductById retrieves a single product by its ID, including its variants
func GetProductById(c *gin.Context) {
	var product models.Product
	productId := c.Param("id")

	if err := config.DB.Preload("Variants").First(&product, productId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product with variants
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Save the product and its variants
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct updates an existing product by its ID
func UpdateProduct(c *gin.Context) {
	var product models.Product
	productId := c.Param("id")

	// Check if the product exists
	if err := config.DB.First(&product, productId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Bind the request data to the product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Save the updated product
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product by its ID
func DeleteProduct(c *gin.Context) {
	productId := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, productId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
