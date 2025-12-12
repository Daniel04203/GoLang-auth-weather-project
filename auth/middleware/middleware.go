package middleware

import (
	"strings"

	"example.com/labwork_8/auth/model"
	"github.com/gofiber/fiber/v2"
)

func CheckCreds(c *fiber.Ctx) error {
	var req model.User

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query params!",
		})
	}

	if strings.TrimSpace(req.Login) == "" || strings.TrimSpace(req.Password) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Empty fields!",
		})
	}

	return c.Next()
}

func CheckTokenBasic(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing authorization header!",
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or malformed authorization header!",
		})
	}

	if strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer ")) == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or malformed authorization header!",
		})
	}

	return c.Next()
}
