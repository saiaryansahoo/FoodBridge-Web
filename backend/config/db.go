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

	// Auto-migrate tables
	DB.AutoMigrate(&models.User{}, &models.Food{})
	log.Println("Database connected & migrated successfully.")
}
