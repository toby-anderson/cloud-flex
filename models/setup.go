package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var db *gorm.DB

func ConnectDataBase() {

	godotErr := godotenv.Load(".env")

	if godotErr != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"))

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to the database")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("Database connection successful")
	}

	db.AutoMigrate(&User{})

}
