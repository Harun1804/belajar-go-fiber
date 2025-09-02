package routes

import (
	"belajar-go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	users := r.Group("/users")
	users.Get("/", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUserById)
	users.Post("/", controllers.CreateUsers)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)
}
