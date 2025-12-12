package handlers

import (
	"example.com/labwork_8/weather/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/forecast/now",
		middleware.CheckAPIKeyValid,
		middleware.CheckLatLong,
		ForecastNowHandler)

	api.Get("/forecast/history",
		middleware.CheckDates,
		ForecastHistoryHandler)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Route not found!",
		})
	})
}
