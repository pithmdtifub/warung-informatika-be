package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
			errors[err.Field()] = "Error on " + err.Field() + ": " + err.Tag()
			fmt.Print(err.StructField())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	userDB, err := repo.GetUserByUsername(userReq.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Failed to login", "error": err})
	}

	if userDB.Username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": "incorrect username or password"})
	}

	if !helpers.VerifyPassword(userReq.Password, userDB.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": "incorrect username or password"})
	}

	token, err := helpers.GenerateJWT(userReq.Username, userDB.Role)
	return c.JSON(fiber.Map{"message": "Login success", "token": token})
}
