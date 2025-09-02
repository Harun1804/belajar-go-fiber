package models

import "time"

type Book struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
