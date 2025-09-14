package controllers

import (
	"belajar-go-fiber/modules/auth/dtos"
	"belajar-go-fiber/modules/auth/services"
	"belajar-go-fiber/modules/auth/validators"
	"belajar-go-fiber/utils/responseformatter"

	"github.com/gofiber/fiber/v2"
)

var authService = services.NewAuthService()

func Login(c *fiber.Ctx) error {
	user := new(dtos.LoginRequest)

	if err := c.BodyParser(user); err != nil {
		return responseformatter.SendError(c, fiber.StatusServiceUnavailable, "Failed to parse body", err.Error())
	}

	if messages, errValidate := validators.ValidateLoginRequest(user); errValidate != nil {
		return responseformatter.SendError(c, fiber.StatusUnprocessableEntity, "Failed to validate", messages)
	}

	authResponse, err := authService.Login(user)
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusUnauthorized, "Failed to login", err.Error())
	}

	return responseformatter.SendSuccess(c, "Login Success", authResponse)
}

func Register(c *fiber.Ctx) error {
	user := new(dtos.RegisterRequest)

	if err := c.BodyParser(user); err != nil {
		return responseformatter.SendError(c, fiber.StatusServiceUnavailable, "Failed to parse body", err.Error())
	}

	if messages, errValidate := validators.ValidateRegisterRequest(user); errValidate != nil {
		return responseformatter.SendError(c, fiber.StatusUnprocessableEntity, "Failed to validate", messages)
	}

	if err := authService.Register(user); err != nil {
		return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}

	return responseformatter.SendSuccess(c, "Register Success")
}