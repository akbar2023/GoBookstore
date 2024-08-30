package database

import (
	"example/web-service-gin/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitTestDB initializes a test database connection
func InitTestDB() *gorm.DB {
	// Retrieve test database credentials from environment variables
	dbUser := os.Getenv("root")
	dbPassword := os.Getenv("")
	dbHost := os.Getenv("127.0.0.1")
	dbPort := os.Getenv("3306")
	dbName := os.Getenv("bookstore")

	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to the test database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to test database: %v", err)
	}

	// Migrate the schema for testing purposes
	err = db.AutoMigrate(&models.Book{}) // Migrate only test models
	if err != nil {
		log.Fatalf("Error migrating test database: %v", err)
	}

	// Assign the database connection to the global DB variable
	DB = db

	return DB
}

// CloseTestDB closes the test database connection
func CloseTestDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting DB from GORM: %v", err)
	}

	// Close the database connection
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Error closing test database: %v", err)
	}
}
