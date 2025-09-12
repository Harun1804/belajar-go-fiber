package models

import (
	"time"
)

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" form:"title" validate:"required" gorm:"size:100,not null"`
	Author    string    `json:"author" form:"author" validate:"required" gorm:"size:100"`
	Year      int       `json:"year" form:"year" validate:"required,number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}