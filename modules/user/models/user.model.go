package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID        int            `json:"id" gorm:"primaryKey"`
    Name      string         `json:"name" gorm:"not null"`
    Email     string         `json:"email" gorm:"not null, unique"`
    Password  string         `json:"-" gorm:"not null,column:password"`
    Phone     string         `json:"phone"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
