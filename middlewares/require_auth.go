package middlewares

import "github.com/gofiber/fiber/v2"

func RequireAuth(c *fiber.Ctx) error {
	err := c.Next()

	return err
}
