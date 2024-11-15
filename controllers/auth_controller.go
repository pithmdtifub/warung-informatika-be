package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/helpers"
	"warung-informatika-be/models"
	repo "warung-informatika-be/repositories"
)

func Login(c *fiber.Ctx) error {
	var validate = validator.New()

	var user models.User

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse JSON", "error": err})
	}

	if err = validate.Struct(user); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "Error on " + err.Field() + ": " + err.Tag()
			fmt.Print(err.StructField())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errors})
	}

	userDB, _ := repo.GetUserByUsername(user.Username)
	if userDB.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": "incorrect username or password"})
	}

	if !helpers.VerifyPassword(user.Password, userDB.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": "incorrect username or password"})
	}

	token, err := helpers.GenerateJWT(user.Username, userDB.Role)
	return c.JSON(fiber.Map{"message": "login success", "token": token})
}
