package config

import (
	"log"
	"time"

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

	// ✅ Enable WAL mode to prevent database locks
	DB.Exec("PRAGMA journal_mode=WAL;")

	// ✅ Auto-migrate tables
	DB.AutoMigrate(&models.User{}, &models.Food{}, &models.Contact{})
	log.Println("Database connected & migrated successfully.")
}

// ✅ Safe Write Mechanism to Handle "Database is Locked"
func SafeDBWrite(action func() error) error {
	maxRetries := 5
	delay := time.Millisecond * 500

	for i := 0; i < maxRetries; i++ {
		err := action()
		if err == nil {
			return nil
		}
		if err.Error() == "database is locked" {
			log.Println("Database is locked, retrying...")
			time.Sleep(delay)
			continue
		}
		return err
	}
	return nil
}
