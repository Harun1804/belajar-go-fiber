package models

import (
	"time"
)

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"size:100,not null"`
	Author    string    `json:"author" gorm:"size:100"`
	Year      int       `json:"year" gorm:"type:int"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}