package controllers

import (
	"ecommerce/config"
	"ecommerce/error"
	"ecommerce/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOrders(c *gin.Context) {
	log.Println("GetOrders Called")

	var orders []models.Order
	if err := config.DB.Preload("Items").Find(&orders).Error; err != nil {
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to get order", error.INTERNAL_ERROR)
		return
	}
	models.CreateSuccessResponse(c, http.StatusOK, "", orders)
}

func CreateOrder(c *gin.Context) {
	log.Println("CreateOrder Called")

	var request models.OrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusBadRequest, "Invalid input", error.ErrInvalidRequest)
		return
	}

	// Calculate total amount and verify stock
	totalAmount := 0.0
	for _, item := range request.Items {
		var product models.Product
		var variant models.Variant
		if err := config.DB.First(&product, item.Productid).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				models.CreateErrorResponse(c, http.StatusBadRequest, "Product not found", error.PRODUCT_NOT_FOUND)
			} else {
				models.CreateErrorResponse(c, http.StatusInternalServerError, "Database error", error.INTERNAL_ERROR)
			}
			return
		}
		if err := config.DB.First(&variant, item.Variantid).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				models.CreateErrorResponse(c, http.StatusBadRequest, "Variant not found", error.VARIANT_NOT_FOUND)
			} else {
				models.CreateErrorResponse(c, http.StatusInternalServerError, "Database error", error.INTERNAL_ERROR)
			}
			return
		}

		// Check stock
		if variant.Stock < item.Quantity {
			models.CreateErrorResponse(c, http.StatusBadRequest, "Not enough stock for product", error.NOT_ENOUGH_STOCK)
			return
		}

		// Update stock (Decrement stock for each item)
		variant.Stock -= item.Quantity
		if err := config.DB.Save(&variant).Error; err != nil {
			models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to update stock", error.INTERNAL_ERROR)
			return
		}

		// Add price to total amount
		totalAmount += item.Price * float64(item.Quantity)
	}

	// Create Order
	order := models.Order{
		Userid:          request.UserId,
		Totalamount:     float32(totalAmount),
		Shippingaddress: request.ShippingAddress,
		// Status:          "Pending",
	}

	if err := config.DB.Create(&order).Error; err != nil {
		log.Println("error", err)
		models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to create order", error.INTERNAL_ERROR)
		return
	}

	// Create Order Items
	for _, item := range request.Items {
		orderItem := models.OrderItem{
			Orderid:   order.Orderid,
			Productid: item.Productid,
			Variantid: item.Variantid,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		if err := config.DB.Create(&orderItem).Error; err != nil {
			log.Println("error", err)
			models.CreateErrorResponse(c, http.StatusInternalServerError, "Failed to create order item", error.INTERNAL_ERROR)
			return
		}
	}
	models.CreateSuccessResponse(c, http.StatusOK, "Order placed successfully", order.Orderid)
}
