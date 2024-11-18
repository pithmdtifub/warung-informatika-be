package routes

import (
	"github.com/gofiber/fiber/v2"
	c "warung-informatika-be/controllers"
	m "warung-informatika-be/middlewares"
)

func CategoryRoutes(v fiber.Router) {
	api := v.Group("/categories")

	api.Get("/", c.GetCategories)
	api.Get("/:id", c.GetCategory)
	api.Post("/", m.RequireAuth, m.RequireAdmin, m.RequireJSONContent, c.CreateCategory)
	api.Delete("/:id", c.DeleteCategory)
}
