package services

import (
	"belajar-go-fiber/modules/auth/dtos"
	userDto "belajar-go-fiber/modules/user/dtos"
	"belajar-go-fiber/modules/user/services"
	"belajar-go-fiber/utils"
	"errors"
	"time"
)

type AuthService struct{}

var userService = services.NewUserService()

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Register(userReq *dtos.RegisterRequest) error {
	createUser := userDto.UserCreateRequest{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Phone:    userReq.Phone,
		Password: userReq.Password,
	}

	_, err := userService.CreateUser(&createUser)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) Login(userReq *dtos.LoginRequest) (*dtos.AuthResponse, error) {
	user, err := userService.GetUserByEmail(userReq.Email)
	if err != nil {
		return nil, err
	}

	token := ""
	var expiredAt time.Time
	
	if !utils.CheckPasswordHash(userReq.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	} else {
		// generate token
		token, expiredAt = utils.GenerateToken(user.ID)
	}
	
	dtos := &dtos.AuthResponse{
		Token:    token,
		Type:     "Bearer",
		ExpiredAt: expiredAt.Unix(),
	}

	return dtos, nil
}
