package routes

import (
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/controllers"
	"warung-informatika-be/middlewares"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/v1/auth", middlewares.RequireJSONContent)

	api.Post("/login", controllers.Login)
}
