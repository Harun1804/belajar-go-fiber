package migrations

import (
	"belajar-go-fiber/database"
	"belajar-go-fiber/models"
	"fmt"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		panic("Migration failed")
	}

	fmt.Println("Migration completed successfully")
}