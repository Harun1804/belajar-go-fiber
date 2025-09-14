package dtos

type BookCreateRequest struct {
	Title     string    `json:"title" form:"title" binding:"required"`
	Author    string    `json:"author" form:"author" binding:"required"`
	Year      int       `json:"year" form:"year" binding:"required" validate:"number"`
}

type BookUpdateRequest struct {
	Title     string  `json:"title" form:"title" binding:"required"`
	Author    string  `json:"author" form:"author" binding:"required"`
	Year      int     `json:"year" form:"year" binding:"required" validate:"number"`
	ID        int     `json:"-"`
}