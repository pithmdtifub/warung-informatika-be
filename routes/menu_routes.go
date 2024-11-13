package routes

import (
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/controllers"
	"warung-informatika-be/middlewares"
)

func MenuRoutes(app *fiber.App) {
	api := app.Group("api/v1/menus/")

	api.Get("/", controllers.GetMenus)
	api.Get("/:id", controllers.GetMenu)
	api.Post("/", middlewares.RequireJSONContent, controllers.CreateMenu)
	api.Put("/:id", controllers.UpdateMenu)
	api.Delete("/:id", controllers.DeleteMenu)
}
