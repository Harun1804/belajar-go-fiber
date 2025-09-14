package validators

import (
	"belajar-go-fiber/modules/auth/dtos"
	"belajar-go-fiber/utils"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateLoginRequest(req *dtos.LoginRequest) (map[string][]string, error) {
	err := validate.Struct(req)
	if err != nil {
		errFields := err.(validator.ValidationErrors)
		return utils.FormatValidationErrors(errFields), err
	}
	return nil, nil
}

func ValidateRegisterRequest(req *dtos.RegisterRequest) (map[string][]string, error) {
	err := validate.Struct(req)
	if err != nil {
		errFields := err.(validator.ValidationErrors)
		return utils.FormatValidationErrors(errFields), err
	}
	return nil, nil
}
