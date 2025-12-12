package main

import (
	"log"

	"example.com/labwork_8/gateway/cfg"
	"example.com/labwork_8/gateway/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
}

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Setup routes
	handlers.SetupRoutes(app)

	port := cfg.GetProperty("GATEWAY_PORT")
	if port == "" {
		port = "3005"
	}

	log.Printf("Gateway service starting at %s port.\n", port)
	log.Fatal(app.Listen(":" + port))
}
