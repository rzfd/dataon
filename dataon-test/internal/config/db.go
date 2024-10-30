package config

import (
	"fmt"
	"log"
	"os"

	"github.com/rzfd/dataon-test/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error when connecting to database: %s", err.Error())
	}
	if err := db.AutoMigrate(&model.Kain{}, &model.JenisKain{}, &model.KainKualitas{}, &model.Kualitas{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	DB = db
}
