package handlers

import (
	"fmt"
	"strings"

	"example.com/labwork_8/gateway/cfg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func AuthHandler(c *fiber.Ctx) error {
	host := cfg.GetProperty("AUTH_HOST")
	port := cfg.GetProperty("AUTH_PORT")

	body := c.Body()

	c.Request().SetBody(body)

	return proxy.Do(c, fmt.Sprintf("%s%s/api/auth", host, port))
}

func ForecastNowHandler(c *fiber.Ctx) error {
	host := cfg.GetProperty("FORECAST_HOST")
	port := cfg.GetProperty("FORECAST_PORT")

	queries := c.Context().QueryArgs().String()
	if strings.TrimSpace(queries) == "" {
		queries = ""
	}

	return proxy.Do(c, fmt.Sprintf("%s%s/api/forecast/now?%s", host, port, queries))
}

func ForecastHistoryHandler(c *fiber.Ctx) error {
	host := cfg.GetProperty("FORECAST_HOST")
	port := cfg.GetProperty("FORECAST_PORT")

	queries := c.Context().QueryArgs().String()

	return proxy.Do(c, fmt.Sprintf("%s%s/api/forecast/history?%s", host, port, queries))
}
