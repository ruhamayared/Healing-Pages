package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// .env variable
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println("dsn---->:", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// DB.AutoMigrate(&models.Entry{})
}
