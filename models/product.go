package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID   uint      `json:"product_id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	CategoryID  uint      `json:"category_id"`
	Variants    []Variant `json:"variants,omitempty" gorm:"foreignKey:ProductID"`
}
