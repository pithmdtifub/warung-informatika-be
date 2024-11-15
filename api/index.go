package handler

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"warung-informatika-be/database"
	"warung-informatika-be/routes"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	database.ConnectDatabase()
	//database.Migrate()

	app := fiber.New()
	app.Use(cors.New())
	routes.Routes(app)

	return adaptor.FiberApp(app)
}
