package routes

import (
	"belajar-go-fiber/modules/auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoute(r fiber.Router) {
	auth := r.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}