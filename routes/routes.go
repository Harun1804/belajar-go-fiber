package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	api := r.Group("/api")
	// User routes
	RegisterUserRoute(api)
}
