package db

import (
	"log"

	"example.com/labwork_8/weather/cfg"
	"example.com/labwork_8/weather/model"
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
		panic("Failed to connect database")
	}

	log.Println("Connection opened to database")
	DB.AutoMigrate(&model.ForecastDB{})
	log.Println("Database migrated")
}
