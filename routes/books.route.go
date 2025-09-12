package routes

import (
	"belajar-go-fiber/modules/book/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoute(r fiber.Router) {
	books := r.Group("/books")
	books.Get("/", controllers.GetAllBooks)
	books.Get("/:id", controllers.GetBookByID)
	books.Post("/", controllers.CreateBook)
	books.Put("/:id", controllers.UpdateBook)
	books.Delete("/:id", controllers.DeleteBook)
}