package database

import (
	"belajar-go-fiber/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
		// Load konfigurasi database dari .env
	dbUser := configs.GetEnv("DB_USER", "root")
	dbPass := configs.GetEnv("DB_PASS", "")
	dbHost := configs.GetEnv("DB_HOST", "localhost")
	dbPort := configs.GetEnv("DB_PORT", "3306")
	dbName := configs.GetEnv("DB_NAME", "latihan")


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	dbUser, dbPass, dbHost, dbPort, dbName)
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connected")
}