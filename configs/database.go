package configs

import (
	"belajar-go-fiber/modules/books/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "3306")
	dbUser := GetEnv("DB_USER", "root")
	dbPass := GetEnv("DB_PASS", "")
	dbName := GetEnv("DB_NAME", "perpustakaan")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Database connected successfully")

	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}
	fmt.Println("Database migrated successfully")
}