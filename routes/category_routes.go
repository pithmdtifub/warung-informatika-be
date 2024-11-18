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

	api.Use(m.RequireAuth, m.RequireAdmin)

	api.Post("/", m.RequireJSONContent, c.CreateCategory)
	api.Put("/:id", m.RequireJSONContent, c.UpdateCategory)
	api.Delete("/:id", c.DeleteCategory)
}
