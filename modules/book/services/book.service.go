package services

import (
	"belajar-go-fiber/database"
	"belajar-go-fiber/modules/book/dtos"
	"belajar-go-fiber/modules/book/models"
)

type BookService struct{}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) GetAllBooks() ([]*dtos.BookResponse, error) {
	var books []models.Book
	if err := database.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	var bookResponses []*dtos.BookResponse 
	for _, book := range books {
		bookResponses = append(bookResponses, &dtos.BookResponse{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Year:   book.Year,
		})
	}

	return bookResponses, nil
}

func (s *BookService) GetBookById(id int) (*dtos.BookResponse, error) {
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	
	bookResponse := &dtos.BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}
	return bookResponse, nil
}

func (s *BookService) CreateBook(bookReq *dtos.BookCreateRequest) (*models.Book, error) {
	newBook := models.Book{
		Title:  bookReq.Title,
		Author: bookReq.Author,
		Year:   bookReq.Year,
	}

	if err := database.DB.Debug().Create(&newBook).Error; err != nil {
		return nil, err
	}
	return &newBook, nil
}

func (s *BookService) UpdateBook(bookReq *dtos.BookUpdateRequest) (*models.Book, error) {
	updatedBook := models.Book{
		ID:     bookReq.ID,
		Title:  bookReq.Title,
		Author: bookReq.Author,
		Year:   bookReq.Year,
	}

	if err := database.DB.Debug().Model(&models.Book{}).Where("id = ?", bookReq.ID).Updates(updatedBook).Error; err != nil {
		return nil, err
	}

	return &updatedBook, nil
}

func (S *BookService) DeleteBook(id int) error {
	if err := database.DB.Delete(&models.Book{}, id).Error; err != nil {
		return err
	}

	return nil
}