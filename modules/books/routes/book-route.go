package routes

import (
	"belajar-go-fiber/modules/books/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoutes(app fiber.Router) {
	bookGroup := app.Group("/books")
	bookGroup.Get("/", controllers.GetAllBooks)
}