package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"example.com/labwork_8/weather/cfg"
	"example.com/labwork_8/weather/db"
	"example.com/labwork_8/weather/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
}

func main() {
	// Init
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	// Routes
	handlers.SetupRoutes(app)

	// DB
	db.ConnectDB()

	// Fetching data
	duration, err := strconv.Atoi(cfg.GetProperty("WEATHER_FREQ_SECS"))

	if err != nil {
		duration = 3600
	}

	go handlers.ForecastSheduler(time.Duration(duration))

	// Setting up port
	port := cfg.GetProperty("FORECAST_PORT")

	if strings.TrimSpace(port) == "" {
		port = "3007"
	}

	log.Printf("Weather service starting at %s port.\n", port)
	log.Fatal(app.Listen(":" + port))
}
