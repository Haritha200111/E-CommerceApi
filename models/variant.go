package models

type Product struct {
	ProductID           uint   `json:"productId" gorm:"primaryKey;column:product_id"`
	ProductName         string `json:"productName" validate:"required"`
	Description         string `json:"description"`
	ProductCategoryName string `json:"productcategoryname" gorm:"foreignKey:productcategoryname;references:category_name"`
	// Variants            []Variant `gorm:"foreignKey:ProductID;references:category_name"`
}

type UpdateProductInput struct {
	ProductName         string `json:"productName"`
	NewProductName      string `json:"newProductName"`
	Description         string `json:"description"`
	ProductCategoryName string `json:"productCategoryName"`
}

func (Product) TableName() string {
	return "product"
}

type Variant struct {
	VariantID   uint    `json:"variantId" gorm:"primaryKey;column:variantid"`
	Productname string  `json:"productname" validate:"required" gorm:"column:productname"`
	Price       float64 `json:"price" validate:"required"`
	Size        int     `json:"size"`
	Color       string  `json:"color"`
	Stock       int     `json:"stock" validate:"required"`
}

type VariantUpdate struct {
	VariantID uint    `json:"variantId" gorm:"primaryKey;column:variantid"`
	Price     float64 `json:"price" `
	Size      int     `json:"size"`
	Color     string  `json:"color"`
	Stock     int     `json:"stock"`
}
