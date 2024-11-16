package controllers

import (
	"fmt"
	"strconv"
	db "warung-informatika-be/database"
	"warung-informatika-be/models"
	"warung-informatika-be/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetMenus(c *fiber.Ctx) error {
	search := c.Query("search", "")
	categoryID, _ := strconv.Atoi(c.Query("category_id", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	menus, err := repositories.GetMenus(search, categoryID, limit, offset)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get menus", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully get all menu", "data": menus})
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
