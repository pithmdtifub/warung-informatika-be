package routes

import (
	"github.com/gofiber/fiber/v2"
	"warung-informatika-be/middlewares"
)

func Routes(app *fiber.App) {
	app.Use(middlewares.Logger)

	AuthRoutes(app)
	MenuRoutes(app)
	CategoryRoutes(app)
}
