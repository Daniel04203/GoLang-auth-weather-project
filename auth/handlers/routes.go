package handlers

import (
	"example.com/labwork_8/auth/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/verify",
		middleware.CheckTokenBasic,
		VerifyHandler)

	api.Post("/auth",
		middleware.CheckCreds,
		AuthHandler)
}
