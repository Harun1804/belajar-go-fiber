package dtos

type UserCreateRequest struct {
    Name     string `json:"name" binding:"required" validate:"required"`
    Email    string `json:"email" binding:"required" validate:"required,email,unique_email"`
    Phone    string `json:"phone" validate:"omitempty"`
    Password string `json:"password" binding:"required" validate:"gte=5"`
}

type UserUpdateRequest struct {
    Name     string  `json:"name" binding:"required" validate:"required"`
    Email    string  `json:"email" binding:"required" validate:"required,email,unique_email"`
    Phone    string  `json:"phone" validate:"omitempty"`
    Password *string `json:"password,omitempty" validate:"omitempty,gte=5"`
    ID       int     `json:"-"`
}
