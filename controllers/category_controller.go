package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/models"
	repo "warung-informatika-be/repositories"
)

func GetCategories(c *fiber.Ctx) error {
	categories, err := repo.GetCategories()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all category", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully get all category", "data": categories})
}

func GetCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	category, _ := repo.GetCategory(id)

	if category.ID < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found", "error": "category not found"})
	}

	return c.JSON(fiber.Map{"message": "Successfully get category", "data": category})
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

	if err := repo.CreateCategory(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Category created successfully", "data": category})
}
