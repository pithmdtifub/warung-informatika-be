package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func RequireAdmin(c *fiber.Ctx) error {
	key := os.Getenv("JWT_SECRET")
	var jwtKey = []byte(key)

	tokenString := c.Get("Authorization")
	tokenArray := strings.Split(tokenString, " ")
	if len(tokenArray) == 1 {
		tokenString = tokenArray[0]
	} else if len(tokenArray) == 2 {
		tokenString = tokenArray[1]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	claims := token.Claims.(jwt.MapClaims)

	if claims["role"] != "Admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized", "error": "You are not authorized to access this resource"})
	}

	err = c.Next()

	return err
}
