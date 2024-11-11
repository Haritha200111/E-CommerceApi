package models

// type Category struct {
// 	CategoryID   int16  `json:"categoryId" gorm:"primaryKey;column:category_id"` // Ensure it's a positive number
// 	CategoryName string `json:"categoryName,omitempty" validate:"required,gt=0"` // Ensure it's positive if provided
// 	// SubCategoryName    pq.StringArray `json:"subcategoryName" gorm:"type:text[]" validate:"required"` // Ensure it's not empty and at least 3 characters
// }

// func (Category) TableName() string {
// 	return "category"
// }

// type SubCategory struct {
// 	SubCategoryID      uint   `gorm:"primaryKey;column:sub_category_id" json:"sub_category_id"`
// 	ParentCategoryName string `json:"parent_category_name"`
// 	SubCategoryName    string `json:"sub_category_name"` // Changed to string since we handle one at a time
// }

// func (SubCategory) TableName() string {
// 	return "subcategory"
// }

type Category struct {
	CategoryID    int16         `json:"categoryId" gorm:"primaryKey;column:category_id"`
	CategoryName  string        `json:"categoryName,omitempty" validate:"required,gt=0"`
	SubCategories []SubCategory `gorm:"foreignKey:parent_category_name;references:category_name" json:"subcategories,omitempty"`
}

func (Category) TableName() string {
	return "category"
}

type SubCategory struct {
	SubCategoryID      uint   `gorm:"primaryKey;column:sub_category_id" json:"sub_category_id"`
	ParentCategoryName string `json:"parent_category_name"`
	SubCategoryName    string `json:"sub_category_name"`
}

func (SubCategory) TableName() string {
	return "subcategory"
}

type UpdateCategoryInput struct {
	CategoryName       string `json:"categoryName" validate:"required,gt=0"`
	UpdateCategoryName string `json:"updatecategoryName" validate:"required,gt=0"`
}
