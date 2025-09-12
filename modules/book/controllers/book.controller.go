package controllers

import (
	"belajar-go-fiber/database"
	"belajar-go-fiber/modules/book/models"
	"belajar-go-fiber/utils/responseformatter"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	var books []models.Book

	database.DB.Find(&books)

	return c.Status(fiber.StatusOK).JSON(responseformatter.SuccessResponse{
		Status:  true,
		Message: "Success Get All Books",
		Data:    books,
	})
}

func GetBookByID(c *fiber.Ctx) error {
	var book []*models.Book

	result := database.DB.Find(&book, "id = ?", c.Params("id"))

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responseformatter.ErrorResponse{
			Status:  false,
			Message: "Failed to retrieve book",
			Data:    result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responseformatter.SuccessResponse{
		Status:  true,
		Message: "Success Get Book",
		Data:    book,
	})
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(responseformatter.ErrorResponse{
			Status:  false,
			Message: "Failed to parse body",
			Data:    err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responseformatter.ErrorResponse{
			Status:  false,
			Message: "Failed to validate",
			Data:    errValidate.Error(),
		})
	}

	newBook := models.Book{
		Title:       book.Title,
		Author:      book.Author,
		Year:        book.Year,
	}

	database.DB.Debug().Create(&newBook)

	return c.Status(fiber.StatusOK).JSON(responseformatter.SuccessResponse{
		Status:  true,
		Message: "Success Created new Book",
	})
}

func UpdateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(responseformatter.ErrorResponse{
			Status:  false,
			Message: "Failed to parse body",
			Data:    err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responseformatter.ErrorResponse{
			Status:  false,
			Message: "Failed to validate",
			Data:    errValidate.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Debug().Model(&book).Where("id = ?", id).Updates(models.Book{
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	})

	return c.Status(fiber.StatusOK).JSON(responseformatter.SuccessResponse{
		Status:  true,
		Message: "Success Updated Book",
	})
}

func DeleteBook(c *fiber.Ctx) error {
	var book models.Book
	result := database.DB.Delete(&book, c.Params("id"))

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responseformatter.ErrorResponse{
			Status:  false,
			Message: "Failed to delete book",
			Data:    result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responseformatter.SuccessResponse{
		Status:  true,
		Message: "Success Deleted Book",
	})
}