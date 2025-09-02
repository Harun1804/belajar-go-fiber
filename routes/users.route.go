package routes

import (
	"belajar-go-fiber/modules/user/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoute(r fiber.Router) {
	users := r.Group("/users")
	users.Get("/", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUserById)
	users.Post("/", controllers.CreateUsers)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)
}