package routes

import (
	"github.com/gofiber/fiber/v2"
	c "warung-informatika-be/controllers"
	m "warung-informatika-be/middlewares"
)

func AuthRoutes(v fiber.Router) {
	api := v.Group("/auth", m.RequireJSONContent)

	api.Post("/login", c.Login)
}
