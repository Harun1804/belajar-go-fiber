package validators

import (
	"belajar-go-fiber/database"
	"belajar-go-fiber/modules/user/dtos"
	"belajar-go-fiber/modules/user/models"

	"github.com/go-playground/validator/v10"
)

// UniqueEmail is a custom validator for unique email
func UniqueEmail(fl validator.FieldLevel) bool {
    email := fl.Field().String()
    var user models.User
    result := database.DB.Where("email = ?", email).First(&user)
    if result.RowsAffected == 0 {
        return true
    }
    // Try to get ID from parent struct
    parent := fl.Parent().Interface()
    switch v := parent.(type) {
    case dtos.UserUpdateRequest:
        return user.ID == v.ID
    default:
        // For create, any found user means not unique
        return false
    }
}
