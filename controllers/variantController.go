package controllers

import (
	"ecommerce/config"
	"ecommerce/error"
	"ecommerce/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVariants(c *gin.Context) {
	log.Println("GetVariants Called")

	var variants []models.Variant

	// Retrieve the variants for a specific product by product ID
	if err := config.DB.Find(&variants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, variants)
}

func GetVariantById(c *gin.Context) {
	log.Println("GetVariantById Called")

	var input models.Variant
	// Bind incoming JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}
	var variant models.Variant
	// Fetch the variant by ID
	if err := config.DB.Where("productname = ?", input.Productname).First(&variant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "variant not found"})
		return
	}
	c.JSON(http.StatusOK, variant)
}

func CreateVariant(c *gin.Context) {
	log.Println("CreateVariant Called")

	var input models.Variant

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	if err := validate.Struct(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Failed in validating the input", error.ErrInvalidRequest)
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create variant"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Variant created successfully"})
}

func UpdateVariant(c *gin.Context) {
	log.Println("UpdateVariant Called")

	var input models.VariantUpdate

	// Bind the input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Find the existing product
	var variant models.Variant
	if err := config.DB.Where("variantid = ?", input.VariantID).First(&variant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "variant not found"})
		return
	}

	// Prepare the update data
	updateData := map[string]interface{}{}
	if input.Size != 0 {
		updateData["size"] = input.Size
	}
	if input.Color != "" {
		updateData["color"] = input.Color
	}
	if input.Price != 0 {
		updateData["price"] = input.Price
	}
	if input.Stock != 0 {
		updateData["stock"] = input.Stock
	}

	// Perform the update using the Model and Updates methods
	if len(updateData) > 0 {
		if err := config.DB.Model(&variant).Where("variantid = ?", input.VariantID).Updates(updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update variant"})
			return
		}
	}

	// Return the updated product
	c.JSON(http.StatusOK, variant)
}

func DeleteVariant(c *gin.Context) {
	log.Println("DeleteVariant Called")

	var input models.Variant

	// Bind incoming JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Use Where clause to delete category by CategoryName and its subcategories
	if err := config.DB.Where("variantid = ?", input.VariantID).Delete(&models.Variant{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Variant deleted"})
}
