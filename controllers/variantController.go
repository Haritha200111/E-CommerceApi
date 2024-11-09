package controllers

import (
	"E-COMMERCEAPI/config"
	"E-COMMERCEAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetVariants retrieves all variants for a given product ID
func GetVariants(c *gin.Context) {
	var variants []models.Variant
	productId := c.Param("product_id")

	// Retrieve the variants for a specific product by product ID
	if err := config.DB.Where("product_id = ?", productId).Find(&variants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, variants)
}

// GetVariantById retrieves a single variant by its ID
func GetVariantById(c *gin.Context) {
	var variant models.Variant
	variantId := c.Param("id")

	// Fetch the variant by ID
	if err := config.DB.First(&variant, variantId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Variant not found"})
		return
	}
	c.JSON(http.StatusOK, variant)
}

// CreateVariant creates a new variant for a given product
func CreateVariant(c *gin.Context) {
	var variant models.Variant
	if err := c.ShouldBindJSON(&variant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure product exists before creating the variant
	var product models.Product
	if err := config.DB.First(&product, variant.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Save the variant
	if err := config.DB.Create(&variant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create variant"})
		return
	}
	c.JSON(http.StatusCreated, variant)
}

// UpdateVariant updates an existing variant by its ID
func UpdateVariant(c *gin.Context) {
	var variant models.Variant
	variantId := c.Param("id")

	// Check if the variant exists
	if err := config.DB.First(&variant, variantId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Variant not found"})
		return
	}

	// Bind the request data to the variant
	if err := c.ShouldBindJSON(&variant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Save the updated variant
	if err := config.DB.Save(&variant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update variant"})
		return
	}
	c.JSON(http.StatusOK, variant)
}

// DeleteVariant deletes a variant by its ID
func DeleteVariant(c *gin.Context) {
	variantId := c.Param("id")
	if err := config.DB.Delete(&models.Variant{}, variantId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete variant"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Variant deleted successfully"})
}
