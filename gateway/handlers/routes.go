package handlers

import (
	"example.com/labwork_8/gateway/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth",
		AuthHandler)

	api.Get("/forecast/now",
		middleware.JWTAuthCheck,
		ForecastNowHandler)

	api.Get("/forecast/history",
		middleware.JWTAuthCheck,
		ForecastHistoryHandler)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Route not found!",
		})
	})
}
