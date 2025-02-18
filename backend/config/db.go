package config

import (
	"log"

	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("foodbridge.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// Auto-migrate tables (including Contact)
	DB.AutoMigrate(&models.User{}, &models.Food{}, &models.Contact{})
	log.Println("Database connected & migrated successfully.")
}
