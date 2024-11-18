package routes

import (
	"github.com/gofiber/fiber/v2"
	c "warung-informatika-be/controllers"
	m "warung-informatika-be/middlewares"
)

func MenuRoutes(v fiber.Router) {
	api := v.Group("/menus")

	api.Get("/", c.GetMenus)
	api.Get("/:id", c.GetMenu)

	api.Use(m.RequireAuth, m.RequireAdmin)

	api.Post("/", m.RequireJSONContent, c.CreateMenu)
	api.Put("/:id", c.UpdateMenu)
	api.Delete("/:id", c.DeleteMenu)
}
