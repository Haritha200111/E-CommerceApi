package controllers

import (
	"ecommerce/config"
	"ecommerce/error"
	"ecommerce/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	log.Println("GetProducts Called")

	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	log.Println("GetProductByID Called")

	var input models.Product

	// Bind incoming JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}
	var product models.Product
	// Fetch category by CategoryName
	if err := config.DB.Where("product_name = ?", input.ProductName).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	log.Println("CreateProduct Called")

	var product models.Product

	// Bind incoming JSON to product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Validate the input
	if err := validate.Struct(&product); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Failed in validating the input", error.ErrInvalidRequest)
		return
	}

	// Save to database
	if err := config.DB.Create(&product).Error; err != nil {
		log.Println("Error while saving product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func UpdateProduct(c *gin.Context) {
	log.Println("UpdateProduct Called")

	var input models.UpdateProductInput

	// Bind the input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Find the existing product
	var product models.Product
	if err := config.DB.Where("product_name = ?", input.ProductName).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Prepare the update data
	updateData := map[string]interface{}{}
	if input.NewProductName != "" {
		updateData["product_name"] = input.NewProductName
	}
	if input.Description != "" {
		updateData["description"] = input.Description
	}
	if input.ProductCategoryName != "" {
		updateData["product_category_name"] = input.ProductCategoryName
	}

	// Perform the update using the Model and Updates methods
	if len(updateData) > 0 {
		if err := config.DB.Model(&product).Where("product_name = ?", input.ProductName).Updates(updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}
	}

	// Return the updated product
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	log.Println("DeleteProduct Called")

	var input models.Product

	// Bind incoming JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Use Where clause to delete category by CategoryName and its subcategories
	if err := config.DB.Where("product_name = ?", input.ProductName).Delete(&models.Product{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
