package controllers

import (
	"E-COMMERCEAPI/config"
	"E-COMMERCEAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrders retrieves all orders with their items
func GetOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrderById retrieves a single order by its ID
func GetOrderById(c *gin.Context) {
	var order models.Order
	orderId := c.Param("id")

	if err := config.DB.Preload("Items").First(&order, orderId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// CreateOrder creates a new order and manages inventory
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Begin transaction for order creation and inventory update
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	orderTotal := 0.0

	// Process each item in the order
	for _, item := range order.Items {
		var variant models.Variant

		// Check if variant exists
		if err := tx.First(&variant, item.VariantID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"error": "Variant not found"})
			return
		}

		// Check inventory availability
		if variant.Quantity < item.Quantity {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for variant"})
			return
		}

		// Update inventory quantity
		variant.Quantity -= item.Quantity
		if err := tx.Save(&variant).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
			return
		}

		// Calculate item price based on quantity and variant's price
		itemPrice := float64(item.Quantity) * variant.MRP
		item.Price = itemPrice
		orderTotal += itemPrice
	}

	// Set order total and status
	order.OrderTotal = orderTotal
	order.Status = "Accepted"

	// Save the order and associated items in the database
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, order)
}

// DeleteOrder deletes an order by its ID
func DeleteOrder(c *gin.Context) {
	orderId := c.Param("id")
	if err := config.DB.Delete(&models.Order{}, orderId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
