package middleware

import (
	"strings"
	"time"

	"example.com/labwork_8/weather/cfg"
	"example.com/labwork_8/weather/model"
	owm "github.com/briandowns/openweathermap"
	"github.com/gofiber/fiber/v2"
)

func CheckAPIKeyValid(c *fiber.Ctx) error {
	key := cfg.GetProperty("WEATHER_API_KEY")

	err := owm.ValidAPIKey(key)

	if err != nil || strings.TrimSpace(key) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Incorrect or empty API key!",
		})
	}

	return c.Next()
}

func CheckLatLong(c *fiber.Ctx) error {
	var req model.WeatherNowReq

	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query params!",
		})
	}

	if !(-90 < req.Latitude && req.Latitude < 90) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Latitude value is out of range!",
		})
	}

	if !(-180 < req.Longitude && req.Longitude < 180) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Longtitude value is out of range!",
		})
	}

	c.Locals("latitude", req.Latitude)
	c.Locals("longtitude", req.Longitude)

	return c.Next()
}

func CheckDates(c *fiber.Ctx) error {
	var req model.WeatherHistoryReq

	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query params!",
		})
	}

	if req.DateFrom == "" || req.DateTo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Date params are not set!",
		})
	}

	convedDateFrom, errFrom := time.Parse(time.DateOnly, req.DateFrom)
	if errFrom != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid date_from format!",
		})
	}

	convedDateTo, errTo := time.Parse(time.DateOnly, req.DateTo)
	if errTo != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid date_to format!",
		})
	}

	if convedDateFrom.After(convedDateTo) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Date from param is later than date to!",
		})
	}

	c.Locals("dateFrom", req.DateFrom)
	c.Locals("dateTo", req.DateTo)

	return c.Next()
}
