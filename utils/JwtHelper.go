package utils

import (
	"belajar-go-fiber/configs"
	"time"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(configs.GetEnv("JWT_SECRET", "698aa2a142737b1ece9054c4466010dfb075dc4000b11ec82b60681dacf497185bbcd63feb85c158eee29f88b639fedee7aa53a9b24bffb85a85563d79bf3375"))

func GenerateToken(userId int) (string, time.Time) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userId),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return err.Error(), time.Time{}
	}

	return tokenString, expirationTime
}