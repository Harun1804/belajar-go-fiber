package middlewares

import (
	"belajar-go-fiber/configs"
	"belajar-go-fiber/utils/responseformatter"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(configs.GetEnv("JWT_SECRET", "698aa2a142737b1ece9054c4466010dfb075dc4000b11ec82b60681dacf497185bbcd63feb85c158eee29f88b639fedee7aa53a9b24bffb85a85563d79bf3375"))

// RequireAuthHeader is a middleware that checks for the Authorization header
func RequireAuthHeader(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responseformatter.ErrorResponse{
			Status: false,
			Message: "Authorization header is required",
		})
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(responseformatter.ErrorResponse{
			Status: false,
			Message: "Invalid token",
		})
	}

	if claims.Subject == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responseformatter.ErrorResponse{
			Status: false,
			Message: "Token subject is missing",
		})
	}

	c.Set("username", claims.Subject)

	return c.Next()
}


