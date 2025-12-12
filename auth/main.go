package main

import (
	"log"

	"example.com/labwork_8/auth/cfg"
	"example.com/labwork_8/auth/db"
	"example.com/labwork_8/auth/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
}

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	handlers.SetupRoutes(app)

	db.ConnectDB()

	port := cfg.GetProperty("AUTH_PORT")

	if port == "" {
		port = "3006"
	}

	log.Printf("Auth service starting at %s port\n", port)
	log.Fatal(app.Listen(":" + port))
}
