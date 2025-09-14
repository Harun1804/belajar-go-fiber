package validators

import (
	"belajar-go-fiber/modules/user/dtos"
	"belajar-go-fiber/utils"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("unique_email", UniqueEmail)
}

func ValidateUserCreateRequest(req *dtos.UserCreateRequest) (map[string][]string, error) {
	err := validate.Struct(req)
	if err != nil {
		errFields := err.(validator.ValidationErrors)
		return utils.FormatValidationErrors(errFields), err
	}
	return nil, nil
}

func ValidateUserUpdateRequest(req *dtos.UserUpdateRequest) (map[string][]string, error) {
	err := validate.Struct(req)
	if err != nil {
		errFields := err.(validator.ValidationErrors)
		return utils.FormatValidationErrors(errFields), err
	}
	return nil, nil
}
