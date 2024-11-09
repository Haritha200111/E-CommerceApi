package models

import "gorm.io/gorm"

type Variant struct {
	gorm.Model
	VariantID     uint    `json:"variant_id" gorm:"primaryKey"`
	ProductID     uint    `json:"product_id"`
	Name          string  `json:"name"`
	MRP           float64 `json:"mrp"`
	DiscountPrice float64 `json:"discount_price"`
	Size          string  `json:"size"`
	Color         string  `json:"color"`
	Quantity      int     `json:"quantity"`
}
