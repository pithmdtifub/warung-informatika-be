package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"warung-informatika-be/models"
)

func RequireAdmin(c *fiber.Ctx) error {
	claims, ok := c.Locals("claims").(jwt.MapClaims)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error", "error": "Error while getting claims"})
	}

	if claims["role"] != models.RoleAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden", "error": "You are not allowed to access this resource"})
	}

	err := c.Next()

	return err
}
