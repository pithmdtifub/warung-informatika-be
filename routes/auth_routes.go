package routes

import (
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/controllers"
	"warung-informatika-be/middlewares"
)

func AuthRoutes(v fiber.Router) {
	api := v.Group("/auth", middlewares.RequireJSONContent)

	api.Post("/login", controllers.Login)
}
