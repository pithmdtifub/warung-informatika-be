package controllers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
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
		categoriesRes = append(categoriesRes, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return c.JSON(fiber.Map{"message": "Successfully get all category", "data": categoriesRes})
}

func GetCategory(c *fiber.Ctx) error {
	param := struct {
		ID uint `params:"id"`
	}{}
	err := c.ParamsParser(&param)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to get category", "error": "category not found"})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to get category", "error": "invalid category id"})
	}

	category, err := repo.GetCategory(param.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to get category", "error": "category not found"})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get category", "error": err.Error()})
	}

	categoryRes := dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return c.JSON(fiber.Map{"message": "Successfully get category", "data": categoryRes})
}

func CreateCategory(c *fiber.Ctx) error {
	validate := validator.New()

	categoryReq := new(dto.CategoryRequest)

	if err := c.BodyParser(&categoryReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err.Error()})
	}

	if err := validate.Struct(categoryReq); err != nil {
		_errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			_errors[field] = "Error on " + field + ": " + err.Tag()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": _errors})
	}

	category := models.Category{Name: categoryReq.Name}

	if err := repo.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category", "error": err.Error()})
	}

	categoryRes := dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return c.JSON(fiber.Map{"message": "Category created successfully", "data": categoryRes})
}

func UpdateCategory(c *fiber.Ctx) error {
	param := struct {
		ID uint `params:"id"`
	}{}
	err := c.ParamsParser(&param)

	validate := validator.New()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to update category", "error": "category not found"})
	}

	category, err := repo.GetCategory(param.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to update category", "error": "category not found"})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update category", "error": err.Error()})
	}

	categoryReq := new(dto.CategoryUpdateRequest)

	if err := c.BodyParser(categoryReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err.Error()})
	}

	if err := validate.Struct(categoryReq); err != nil {
		_errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			_errors[field] = "Error on " + field + ": " + err.Tag()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": _errors})
	}

	category.Name = categoryReq.Name

	if err := repo.UpdateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update category", "error": err.Error()})
	}

	categoryRes := dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return c.JSON(fiber.Map{"message": "Category updated successfully", "data": categoryRes})
}

func DeleteCategory(c *fiber.Ctx) error {
	param := struct {
		ID uint `params:"id"`
	}{}
	err := c.ParamsParser(&param)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to delete category", "error": "category not found"})
	}

	err = repo.DeleteCategory(param.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete category", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}
