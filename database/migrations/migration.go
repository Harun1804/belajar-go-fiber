package migrations

import (
	"belajar-go-fiber/database"
	userModels "belajar-go-fiber/modules/user/models"
	bookModels "belajar-go-fiber/modules/book/models"
	"fmt"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&userModels.User{},
		&bookModels.Book{},
	)

	if err != nil {
		panic("Migration failed")
	}

	fmt.Println("Migration completed successfully")
}