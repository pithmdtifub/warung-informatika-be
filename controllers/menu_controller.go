package controllers

import (
	"errors"
	"warung-informatika-be/dto"
	"warung-informatika-be/models"
	"warung-informatika-be/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	menusRes := make([]dto.MenuResponse, 0)

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
	param := struct {
		ID uint `params:"id"`
	}{}
	err := c.ParamsParser(&param)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to get menu", "error": "menu not found"})
	}

	menu, err := repositories.GetMenu(param.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Menu not found", "error": "menu not found"})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get menu", "error": err.Error()})
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
		Name:         menuDTO.Name,
		CategoryID:   menuDTO.CategoryID,
		CategoryName: menuDTO.CategoryName,
		Description:  menuDTO.Description,
		Price:        menuDTO.Price,
		Image:        menuDTO.Image,
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
		CategoryName: menu.CategoryName,
		Image:        menu.Image,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Menu created successfully", "data": menuRes})
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
		ID:           params.ID,
		Name:         menuDTO.Name,
		CategoryID:   menuDTO.CategoryID,
		CategoryName: menuDTO.CategoryName,
		Description:  menuDTO.Description,
		Price:        menuDTO.Price,
		Image:        menuDTO.Image,
	}

	err := repositories.UpdateMenu(&menu)

	if err.Error() == "Data not exist" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to update menu", "error": err.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update menu", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Menu updated successfully", "data": menu})
}

func DeleteMenu(c *fiber.Ctx) error {
	var params dto.MenuParams
	if err := c.ParamsParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid menu ID", "error": err.Error()})
	}

	err := repositories.DeleteMenu(int(params.ID))

	if err.Error() == "Data not exist" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to delete menu", "error": err.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete menu", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Menu deleted successfully"})
}
