package controllers

import (
	"belajar-go-fiber/modules/book/dtos"
	"belajar-go-fiber/modules/book/services"
	"belajar-go-fiber/modules/book/validators"
	"belajar-go-fiber/utils/responseformatter"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var bookService = services.NewBookService()

func GetAllBooks(c *fiber.Ctx) error {
	books, err := bookService.GetAllBooks()
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to get all books", err.Error())
	}

	return responseformatter.SendSuccess(c, "Success Get All Books", books)
}

func GetBookByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusBadRequest, "Invalid book ID", err.Error())
	}

	book, err := bookService.GetBookById(id)
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusNotFound, "Book not found", err.Error())
	}

	return responseformatter.SendSuccess(c, "Success Get Book By ID", book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(dtos.BookCreateRequest)

	if err := c.BodyParser(book); err != nil {
		return responseformatter.SendError(c, fiber.StatusServiceUnavailable, "Failed to parse body", err.Error())
	}

	 // Parse file field
	cover, err := c.FormFile("cover")
	if err != nil {
			return responseformatter.SendError(c, fiber.StatusBadRequest, "Cover file is required", err.Error())
	}
	book.Cover = cover

	if messages, errValidate := validators.ValidateBookCreateRequest(book); errValidate != nil {
		return responseformatter.SendError(c, fiber.StatusUnprocessableEntity, "Failed to validate", messages)
	}

	if _, err := bookService.CreateBook(book); err != nil {
		return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to create new book", err.Error())
	}

	return responseformatter.SendSuccess(c, "Success Create New Book")
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusBadRequest, "Invalid book ID", err.Error())
	}
	
	book := new(dtos.BookUpdateRequest)
	if err := c.BodyParser(book); err != nil {
		return responseformatter.SendError(c, fiber.StatusServiceUnavailable, "Failed to parse body", err.Error())
	}

	// Parse file field
	cover, err := c.FormFile("cover")
	if err == nil {
			book.Cover = cover
	}

	if message, errValidate := validators.ValidateBookUpdateRequest(book); errValidate != nil {
		return responseformatter.SendError(c, fiber.StatusUnprocessableEntity, "Failed to validate", message)
	}

	checkBook, err := bookService.GetBookById(id)
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusNotFound, "Book not found", err.Error())
	}

	book.ID = checkBook.ID

	if _, err := bookService.UpdateBook(checkBook.ID, book); err != nil {
		return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to update book", err.Error())
	}

	return responseformatter.SendSuccess(c, "Success Update Book")
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusBadRequest, "Invalid book ID", err.Error())
	}

	_, err = bookService.GetBookById(id)
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusNotFound, "Book not found", err.Error())
	}

	err = bookService.DeleteBook(id)
	if err != nil {
		return responseformatter.SendError(c, fiber.StatusInternalServerError, "Failed to delete book", err.Error())
	}

	return responseformatter.SendSuccess(c, "Success Delete Book")
}
