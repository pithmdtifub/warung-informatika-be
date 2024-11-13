package middlewares

import "github.com/gofiber/fiber/v2"

func RequireJSONContent(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Content type must be application/json"})
	}

	err := c.Next()

	return err
}
