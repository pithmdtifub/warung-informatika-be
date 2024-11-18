package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/dto"
	"warung-informatika-be/models"
	repo "warung-informatika-be/repositories"
)

func GetCategories(c *fiber.Ctx) error {
	categories, err := repo.GetCategories()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all category", "error": err.Error()})
	}

	categoriesRes := make([]dto.CategoryResponse, 0)

	for _, category := range categories {
		menusRes := make([]dto.MenuResponse, 0)
		menus, _ := repo.GetMenusByCategoryId(int(category.ID))

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

		categoriesRes = append(categoriesRes, dto.CategoryResponse{
			ID:    category.ID,
			Name:  category.Name,
			Menus: menusRes,
		})
	}

	return c.JSON(fiber.Map{"message": "Successfully get all category", "data": categoriesRes})
}

func GetCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	category, err := repo.GetCategory(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all category", "error": err.Error()})
	}

	if category.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found", "error": "category not found"})
	}

	menuRes := make([]dto.MenuResponse, 0)
	menus, err := repo.GetMenusByCategoryId(int(category.ID))

	for _, menu := range menus {
		menuRes = append(menuRes, dto.MenuResponse{
			ID:           menu.ID,
			Name:         menu.Name,
			Description:  menu.Description,
			Price:        menu.Price,
			CategoryID:   menu.CategoryID,
			CategoryName: menu.Category.Name,
			Image:        menu.Image,
		})
	}

	categoryRes := dto.CategoryResponse{
		ID:    category.ID,
		Name:  category.Name,
		Menus: menuRes,
	}

	return c.JSON(fiber.Map{"message": "Successfully get category", "data": categoryRes})
}

func CreateCategory(c *fiber.Ctx) error {
	validate := validator.New()

	var categoryReq dto.CategoryRequest

	if err := c.BodyParser(&categoryReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err.Error()})
	}

	if err := validate.Struct(categoryReq); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Error on " + err.Field() + ": " + err.Tag()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	category := models.Category{Name: categoryReq.Name}

	if err := repo.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category", "error": err.Error()})
	}

	categoryRes := dto.CategoryResponse{
		ID:    category.ID,
		Name:  category.Name,
		Menus: make([]dto.MenuResponse, 0),
	}

	return c.JSON(fiber.Map{"message": "Category created successfully", "data": categoryRes})
}
