package controllers

import (
	"categories-service/config"
	"categories-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategories fetches all categories including subcategories
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Preload("Subcategories").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory creates a new category
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategoryByID fetches a single category by ID, including subcategories
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := config.DB.Preload("Subcategories").First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates an existing category
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory deletes a category by ID
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
