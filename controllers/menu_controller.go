package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	db "warung-informatika-be/database"
	"warung-informatika-be/dto"
	"warung-informatika-be/models"
	"warung-informatika-be/repositories"
)

func GetMenus(c *fiber.Ctx) error {
	menus, err := repositories.GetMenus()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get menus", "error": err.Error()})
	}

	var menusRes []dto.MenuResponse
	for _, menu := range menus {
		menusRes = append(menusRes, dto.MenuResponse{
			ID:           menu.ID,
			Name:         menu.Name,
			Description:  menu.Description,
			Price:        menu.Price,
			CategoryID:   menu.CategoryID,
			CategoryName: menu.Category.Name,
			Image:        menu.Image,
		})
	}

	return c.JSON(fiber.Map{"message": "Successfully get all menu", "data": menusRes})
}

func GetMenu(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	menu, err := repositories.GetMenu(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Menu not found", "error": err.Error()})
	}

	menuRes := dto.MenuResponse{
		ID:           menu.ID,
		Name:         menu.Name,
		CategoryID:   menu.CategoryID,
		CategoryName: menu.Category.Name,
		Description:  menu.Description,
		Price:        menu.Price,
		Image:        menu.Image,
	}

	return c.JSON(fiber.Map{"message": "Successfully get menu", "data": menuRes})
}

func CreateMenu(c *fiber.Ctx) error {
	validate := validator.New()

	var menuReq dto.MenuRequest

	err := c.BodyParser(&menuReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err})
	}

	if err = validate.Struct(menuReq); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Error on " + err.Field() + ": " + err.Tag()
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	category := models.Category{ID: menuReq.CategoryID}

	if err := db.DB.First(&category).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid category id", "error": err.Error()})
	}

	menu := models.Menu{
		Name:        menuReq.Name,
		Description: menuReq.Description,
		Price:       menuReq.Price,
		CategoryID:  menuReq.CategoryID,
		Image:       menuReq.Image,
	}

	if err := repositories.CreateMenu(&menu); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create menu", "error": err.Error()})
	}

	menuRes := dto.MenuResponse{
		ID:           menu.ID,
		Name:         menu.Name,
		Description:  menu.Description,
		Price:        menu.Price,
		CategoryID:   menu.CategoryID,
		CategoryName: category.Name,
		Image:        menu.Image,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Menu created successfully", "data": menuRes})
}

func UpdateMenu(c *fiber.Ctx) error {
	return c.Next()
}

func DeleteMenu(c *fiber.Ctx) error {
	return c.Next()
}
