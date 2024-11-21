package main

import (
	"log"
	"os"
	"warung-informatika-be/database"
	"warung-informatika-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDatabase()

	database.Migrate()

	app := fiber.New()
	app.Use(cors.New())
	routes.Routes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
