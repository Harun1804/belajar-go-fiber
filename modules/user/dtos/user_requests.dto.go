package dtos

type UserCreateRequest struct {
    Name     string `json:"name" binding:"required" validate:"required"`
    Email    string `json:"email" binding:"required" validate:"required,email,unique_email"`
    Phone    int    `json:"phone" validate:"omitempty,number"`
    Password string `json:"password" binding:"required" validate:"gte=5"`
}

type UserUpdateRequest struct {
    Name     string  `json:"name" binding:"required" validate:"required"`
    Email    string  `json:"email" binding:"required" validate:"required,email,unique_email"`
    Phone    int     `json:"phone" validate:"omitempty,number,"`
    Password *string `json:"password,omitempty" validate:"omitempty,gte=5"`
    ID       int     `json:"-"`
}
