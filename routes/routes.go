package routes

import (
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/middlewares"
)

func Routes(app *fiber.App) {
	app.Use(middlewares.Logger)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	AuthRoutes(v1)
	MenuRoutes(v1)
	CategoryRoutes(v1)
}
