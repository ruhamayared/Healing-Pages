package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func ConnectDB() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env variable")
	}

	// .env variable
	dsn := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// DB.AutoMigrate(&models.Entry{})
}
