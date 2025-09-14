package validators

import (
	"belajar-go-fiber/modules/book/dtos"
	"belajar-go-fiber/utils"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateBookCreateRequest(req *dtos.BookCreateRequest) (map[string][]string, error) {
	err := validate.Struct(req)
	if err != nil {
		errFields := err.(validator.ValidationErrors)
		return utils.FormatValidationErrors(errFields), err
	}
	return nil, nil
}

func ValidateBookUpdateRequest(req *dtos.BookUpdateRequest) (map[string][]string, error) {
	err := validate.Struct(req)
	if err != nil {
		errFields := err.(validator.ValidationErrors)
		return utils.FormatValidationErrors(errFields), err
	}
	return nil, nil
}