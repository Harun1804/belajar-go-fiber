package main

import (
	"belajar-go-fiber/configs"
	"belajar-go-fiber/database"
	"belajar-go-fiber/database/migrations"
	"belajar-go-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Load environment variables
	configs.LoadEnv()

	// Load cors
	app.Use(cors.New())

	// Load database
	database.InitDB()

	// Migration
	migrations.RunMigration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	routes.RouteInit(app)

	app.Listen(":" + configs.GetEnv("APP_PORT", "4000"))
}