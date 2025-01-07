package middleware

import (
	"github.com/GymLens/Cloud-Computing/config"
	"github.com/gofiber/fiber/v2"
)

func ConfigMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("config", cfg)
		return c.Next()
	}
}
