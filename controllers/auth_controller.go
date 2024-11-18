package controllers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
	"warung-informatika-be/dto"
	"warung-informatika-be/helpers"
	repo "warung-informatika-be/repositories"
)

func Login(c *fiber.Ctx) error {
	validate := validator.New()

	var userReq dto.UserRequest

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err})
	}

	if err := validate.Struct(userReq); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			errors[field] = "Error on " + field + ": " + err.Tag()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	userDB, err := repo.GetUserByUsername(userReq.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": "incorrect username or password"})
	}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Failed to login", "error": err})
	}

	if !helpers.VerifyPassword(userReq.Password, userDB.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": "incorrect username or password"})
	}

	token, err := helpers.GenerateJWT(userReq.Username, userDB.Role)
	return c.JSON(fiber.Map{"message": "Login success", "token": token})
}
