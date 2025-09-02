package main

import (
	"belajar-go-fiber/configs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	configs.LoadEnv()

	// Initialize database
	configs.InitDB()

	app := fiber.New()

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + configs.GetEnv("APP_PORT", "3000"))
}