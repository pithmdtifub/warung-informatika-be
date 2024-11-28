package controllers

import (
	"warung-informatika-be/dto"
	"warung-informatika-be/models"
	"warung-informatika-be/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetMenus(c *fiber.Ctx) error {
	var queryParams dto.MenuQuery
	if err := c.QueryParser(&queryParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query parameters",
			"error":   err.Error(),
		})
	}

	if queryParams.Page < 1 {
		queryParams.Page = 1
	}
	if queryParams.Limit < 1 {
		queryParams.Limit = 10
	}

	menus, err := repositories.GetMenus(queryParams)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get menus", "error": err.Error()})
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
	var menuDTO dto.MenuDTO

	if err := c.BodyParser(&menuDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(menuDTO); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Error on " + err.Field() + ": " + err.Tag()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	menu := models.Menu{
		Name:        menuDTO.Name,
		CategoryID:  menuDTO.CategoryID,
		CategoryName: menuDTO.CategoryName,
		Description: menuDTO.Description,
		Price:       menuDTO.Price,
		Image:       menuDTO.Image,
	}

	if err := repositories.CreateMenu(&menu); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create menu", "error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Menu created successfully", "data": menu})
}

func UpdateMenu(c *fiber.Ctx) error {
	var params dto.MenuParams
	if err := c.ParamsParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid menu ID", "error": err.Error()})
	}

	var menuDTO dto.MenuDTO
	if err := c.BodyParser(&menuDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(menuDTO); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Error on " + err.Field() + ": " + err.Tag()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	menu := models.Menu{
		ID:          params.ID,
		Name:        menuDTO.Name,
		CategoryID:  menuDTO.CategoryID,
		CategoryName: menuDTO.CategoryName,
		Description: menuDTO.Description,
		Price:       menuDTO.Price,
		Image:       menuDTO.Image,
	}

	if err := repositories.UpdateMenu(&menu); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update menu", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Menu updated successfully", "data": menu})
}

func DeleteMenu(c *fiber.Ctx) error {
	var params dto.MenuParams
	if err := c.ParamsParser(&params); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid menu ID", "error": err.Error()})
    }

    _, err := repositories.GetMenu(int(params.ID))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Menu not found", "error": err.Error()})
    }

    if err := repositories.DeleteMenu(int(params.ID)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete menu", "error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "Menu deleted successfully"})
}
