package controllers

import (
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/models"
	"warung-informatika-be/repositories"
)

func GetCategories(c *fiber.Ctx) error {
	categories, err := repositories.GetCategories()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "categories": categories})
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "error": err.Error()})
	}

	if err := repositories.CreateCategory(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "category": category})
}
