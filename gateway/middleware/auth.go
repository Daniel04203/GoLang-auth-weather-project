package middleware

import (
	"fmt"
	"net/http"

	"example.com/labwork_8/gateway/cfg"
	"github.com/gofiber/fiber/v2"
)

func JWTAuthCheck(c *fiber.Ctx) error {
	host := cfg.GetProperty("AUTH_HOST")
	port := cfg.GetProperty("AUTH_PORT")

	authHeader := c.Get("Authorization")

	url := fmt.Sprintf("%s%s/api/auth/verify", host, port)

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to build auth request"})
	}

	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	return c.Next()
}
