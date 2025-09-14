package routes

import (
	"belajar-go-fiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	api := r.Group("/api")
	RegisterAuthRoute(api)

	apiProtected := api.Group("/", middlewares.RequireAuthHeader)
	RegisterUserRoute(apiProtected)
	RegisterBookRoute(apiProtected)
}
