package routes

import (
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/controllers"
	"warung-informatika-be/middlewares"
)

func CategoryRoutes(v fiber.Router) {
	api := v.Group("/categories")

	api.Get("/", controllers.GetCategories)
	api.Post("/", middlewares.RequireJSONContent, middlewares.RequireAuth, controllers.CreateCategory)
}
