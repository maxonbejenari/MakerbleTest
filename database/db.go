package database

import (
	"fmt"
	"github.com/maxonbejenari/testWebApp/config"
	"github.com/maxonbejenari/testWebApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connceting to db %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Patient{})
	DB = db
}
