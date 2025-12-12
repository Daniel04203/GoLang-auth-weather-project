package db

import (
	"log"

	"example.com/labwork_8/auth/cfg"
	"example.com/labwork_8/auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	url := cfg.GetProperty("DB_CONNECTION_URL")

	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		panic("Failed to connect database!")
	}

	log.Println("Connection opened to database")
	DB.AutoMigrate(&model.User{})
	log.Println("Database migrated")
}
