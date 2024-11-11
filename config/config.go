package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// PostgreSQL connection string (adjust the details for your setup)
	dsn := "host=localhost user=postgres password=Monday@01 dbname=E-Commerce port=5432 sslmode=disable"
	var err error
	// Open a connection to the PostgreSQL database using GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
