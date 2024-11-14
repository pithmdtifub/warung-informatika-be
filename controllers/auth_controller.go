package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
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

	userDB, _ := repo.GetUser(nil, user.Username)
	if userDB.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": "incorrect username or password"})
	}

	if !helpers.VerifyPassword(user.Password, userDB.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": "incorrect username or password"})
	}

	token, err := generateJWT(user.Username)
	return c.JSON(fiber.Map{"message": "login success", "token": token})
}

func generateJWT(username string) (string, error) {
	key := os.Getenv("JWT_SECRET")
	var jwtKey = []byte(key)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
