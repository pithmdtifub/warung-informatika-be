package middlewares

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func RequireAuth(c *fiber.Ctx) error {
	key := os.Getenv("JWT_SECRET")
	var jwtKey = []byte(key)

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token not found", "error": "token not found"})
	}

	tokenArray := strings.Split(tokenString, " ")
	if len(tokenArray) == 1 {
		tokenString = tokenArray[1]
	} else if len(tokenArray) == 2 {
		tokenString = tokenArray[0]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token", "error": "invalid token"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token expired", "message": "Token expired"})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token", "error": "invalid token"})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token", "message": "Invalid token"})
	}

	err = c.Next()

	return err
}
