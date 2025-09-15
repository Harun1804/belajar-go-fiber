package dtos

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
    Name     string `json:"name" binding:"required" validate:"required"`
    Email    string `json:"email" binding:"required" validate:"required,email,unique_email"`
    Phone    string `json:"phone" validate:"omitempty"`
    Password string `json:"password" binding:"required" validate:"gte=5"`
}