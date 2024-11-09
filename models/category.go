package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryID    uint       `json:"category_id" gorm:"primaryKey"`
	Name          string     `json:"name" gorm:"not null"`
	Products      []Product  `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
	Subcategories []Category `json:"subcategories,omitempty" gorm:"foreignKey:ParentCategoryID"`
}
