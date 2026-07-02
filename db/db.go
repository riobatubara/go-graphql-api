package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type Book struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"not null"`
	Author     string `gorm:"not null"`
	ISBN       string `gorm:"unique;not null"`
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
}

type Loan struct {
	ID           uint      `gorm:"primaryKey"`
	BookID       uint      `gorm:"not null"`
	Book         Book      `gorm:"foreignKey:BookID"`
	BorrowerName string    `gorm:"not null"`
	LoanDate     time.Time `gorm:"default:CURRENT_DATE"`
	ReturnDate   *time.Time
}

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, fetching operating system contexts")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open standard database streaming:", err)
	}

	fmt.Println("GORM database link successfully hooked up to target container.")
}
