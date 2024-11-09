package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID    uint        `json:"order_id" gorm:"primaryKey"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	OrderTotal float64     `json:"order_total"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	VariantID uint    `json:"variant_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
