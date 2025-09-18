package dtos

import "mime/multipart"

type BookCreateRequest struct {
	Title     string                `json:"title" form:"title" binding:"required"`
	Author    string                `json:"author" form:"author" binding:"required"`
	Year      int                   `json:"year" form:"year" binding:"required" validate:"number"`
	Publisher string                `json:"publisher" form:"publisher" binding:"required"`
	Cover     *multipart.FileHeader `form:"cover" binding:"required"`
}

type BookUpdateRequest struct {
	Title     string                `json:"title" form:"title" binding:"required"`
	Author    string                `json:"author" form:"author" binding:"required"`
	Year      int                   `json:"year" form:"year" binding:"required" validate:"number"`
	Publisher string                `json:"publisher" form:"publisher" binding:"required"`
	Cover     *multipart.FileHeader `form:"cover"`
	ID        int                    `json:"-"`
}