package routes

import (
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/controllers"
	"warung-informatika-be/middlewares"
)

func CategoryRoutes(app *fiber.App) {
	api := app.Group("api/v1/categories")

	api.Get("/", controllers.GetCategories)
	api.Post("/", middlewares.RequireJSONContent, middlewares.RequireAuth, controllers.CreateCategory)
}
