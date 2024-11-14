package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/models"
	"warung-informatika-be/repositories"
)

func GetCategories(c *fiber.Ctx) error {
	categories, err := repositories.GetCategories()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all category", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully get all category", "categories": categories})
}

func CreateCategory(c *fiber.Ctx) error {
	validate := validator.New()

	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err.Error()})
	}

	if err := validate.Struct(category); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Error on " + err.Field() + ": " + err.Tag()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	if err := repositories.CreateCategory(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Category created successfully", "category": category})
}
