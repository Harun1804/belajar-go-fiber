package main

import (
	"belajar-go-fiber/configs"
	"belajar-go-fiber/modules/books/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	configs.LoadEnv()

	// Initialize database
	configs.InitDB()

	app := fiber.New()

	// Register routes
	api := app.Group("/api")
	routes.RegisterBookRoutes(api)

	app.Listen(":" + configs.GetEnv("APP_PORT", "3000"))
}