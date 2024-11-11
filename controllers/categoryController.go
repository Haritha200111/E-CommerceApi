package controllers

import (
	"ecommerce/config"
	"ecommerce/models"
	"log"
	"net/http"

	"ecommerce/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateSubCategory(c *gin.Context) {
	log.Println("CreateSubCategory Called")

	var input struct {
		ParentCategoryName string   `json:"parent_category_name" validate:"required"`
		SubCategoryNames   []string `json:"sub_category_name" validate:"required"`
	}

	// Bind incoming JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Validate the input
	if err := validate.Struct(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Failed in validating the input", error.ErrInvalidRequest)
		return
	}

	// Insert each subcategory name individually
	for _, subCategoryName := range input.SubCategoryNames {
		subCategory := models.SubCategory{
			ParentCategoryName: input.ParentCategoryName,
			SubCategoryName:    subCategoryName,
		}

		// Save to database
		if err := config.DB.Table("subcategory").Omit("created_at", "updated_at", "deleted_at").Create(&subCategory).Error; err != nil {
			log.Println("Error while saving subcategory:", err)
			models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to create subcategory", error.INTERNAL_ERROR)
			return
		}
	}
	models.CreateSuccessResponse(c, http.StatusCreated, "SubCategories created successfully", nil)
}

func CreateCategory(c *gin.Context) {
	log.Println("CreateCategory Called")

	var Category models.Category

	// Bind incoming JSON to subCategory struct
	if err := c.ShouldBindJSON(&Category); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Validate the input
	if err := validate.Struct(&Category); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Failed in validating the input", error.ErrInvalidRequest)
		return
	}

	// Save to database
	if err := config.DB.Table("category").Create(&Category).Error; err != nil {
		log.Println("Error while saving category:", err)
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to create category", error.INTERNAL_ERROR)
		return
	}
	models.CreateSuccessResponse(c, http.StatusCreated, "Category created successfully", nil)
}

func GetCategories(c *gin.Context) {
	log.Println("GetCategories Called")

	var categories []models.Category

	if err := config.DB.Preload("SubCategories").Find(&categories).Error; err != nil {
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to get categories", error.INTERNAL_ERROR)
		return
	}
	models.CreateSuccessResponse(c, http.StatusOK, "", categories)
}

func GetCategoryByID(c *gin.Context) {

	log.Println("GetCategoryByID Called")

	var input models.Category

	// Bind incoming JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Validate the input
	if err := validate.Struct(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Failed in validating the input", error.ErrInvalidRequest)
		return
	}

	var category models.Category
	if err := config.DB.Where("category_name = ?", input.CategoryName).Preload("SubCategories").First(&category).Error; err != nil {
		models.CreateErrorResponse(c, http.StatusNotFound, "Category not found", error.NOT_FOUND_CATEGORY)
		return
	}
	models.CreateSuccessResponse(c, http.StatusOK, "", category)
}

func UpdateCategory(c *gin.Context) {

	log.Println("UpdateCategory Called")

	var input models.UpdateCategoryInput
	var category models.Category
	category.CategoryName = input.UpdateCategoryName

	// Bind incoming JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Validate the input
	if err := validate.Struct(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Failed in validating the input", error.ErrInvalidRequest)
		return
	}

	if err := config.DB.Where("category_name = ?", input.CategoryName).First(&category).Error; err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusNotFound, "Category not found", error.NOT_FOUND_CATEGORY)
		return
	}
	if err := config.DB.Save(&category).Error; err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to update category", error.INTERNAL_ERROR)
		return
	}
	models.CreateSuccessResponse(c, http.StatusOK, "Category updated successfully", category)
}

func DeleteCategory(c *gin.Context) {

	log.Println("DeleteCategory Called")

	var input models.Category

	// Bind incoming JSON to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Validate the input
	if err := validate.Struct(&input); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Failed in validating the input", error.ErrInvalidRequest)
		return
	}

	if err := config.DB.Where("category_name = ?", input.CategoryName).Delete(&models.Category{}).Error; err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed in DB Action", error.INTERNAL_ERROR)
		return
	}
	if err := config.DB.Where("parent_category_name = ?", input.CategoryName).Delete(&models.SubCategory{}).Error; err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed in DB Action", error.INTERNAL_ERROR)
		return
	}
	models.CreateSuccessResponse(c, http.StatusOK, "Category deleted successfully", nil)
}
