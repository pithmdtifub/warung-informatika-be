package controllers

import (
	"fmt"
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
	"warung-informatika-be/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetMenus(c *fiber.Ctx) error {
	search := c.Query("search", "")
	category := c.QueryInt("category", 0)
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	menus, err := repositories.GetMenus(search, category, page, limit)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get menus", "error": err.Error()})
	}

	if len(menus) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No menus found with the given criteria", "data": []models.Menu{}})
	}

	return c.JSON(fiber.Map{"message": "Successfully get all menus", "data": menus})
}

func GetMenu(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	menu, err := repositories.GetMenu(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Menu not found", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully get menu", "menu": menu})
}

func CreateMenu(c *fiber.Ctx) error {
	var validate = validator.New()

	var menu models.Menu

	err := c.BodyParser(&menu)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err})
	}

	if err = validate.Struct(menu); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Error on " + err.Field() + ": " + err.Tag()
			fmt.Print(err.StructField())
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	var category models.Category
	if err := db.DB.First(&category, menu.CategoryID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid category id", "error": err.Error()})
	}

	menu.CategoryName = category.Name

	if err := repositories.CreateMenu(&menu); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create menu", "error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Menu created successfully", "data": menu})
}

func UpdateMenu(c *fiber.Ctx) error {
	return c.Next()
}

func DeleteMenu(c *fiber.Ctx) error {
	return c.Next()
}
