package config

import (
	"ecommerce-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.Variant{}, &models.Order{}, &models.OrderItem{})
	log.Println("Database connected and migrated!")
}
