package services

import (
	"belajar-go-fiber/database"
	"belajar-go-fiber/modules/book/dtos"
	"belajar-go-fiber/modules/book/models"
	"belajar-go-fiber/utils/media"
	"fmt"
)

type BookService struct{}

var pathPrefix = "books/"

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
		coverUrl := ""
		if book.Cover != "" {
			coverUrl, _ = media.GeneratePresignedURL(book.Cover, 3600)
		}

		bookResponses = append(bookResponses, &dtos.BookResponse{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Year:   book.Year,
			Publisher: book.Publisher,
			CoverURL:  coverUrl,
		})
	}

	return bookResponses, nil
}

func (s *BookService) GetBookById(id int) (*dtos.BookResponse, error) {
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return nil, err
	}

	coverUrl := ""
	if book.Cover != "" {
		coverUrl, _ = media.GeneratePresignedURL(book.Cover, 3600)
	}
	
	bookResponse := &dtos.BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
		Publisher: book.Publisher,
		Cover:  book.Cover,
		CoverURL:  coverUrl,
	}

	return bookResponse, nil
}

func (s *BookService) CreateBook(bookReq *dtos.BookCreateRequest) (*models.Book, error) {
	cover := bookReq.Cover
	fileReader, err := cover.Open()
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()
	
	objectName, fileSize, contentType, err := media.ExtractFileData(cover)
	if err != nil {
		return nil, err
	}

	objectName = pathPrefix + objectName
	_, err = media.SendFileToMinio(objectName, fileReader, fileSize, contentType)
	if err != nil {
		return nil, err
	}

	newBook := models.Book{
		Title:  bookReq.Title,
		Author: bookReq.Author,
		Year:   bookReq.Year,
		Publisher: bookReq.Publisher,
		Cover:  objectName,
	}

	if err := database.DB.Debug().Create(&newBook).Error; err != nil {
		return nil, err
	}
	return &newBook, nil
}

func (s *BookService) UpdateBook(id int,bookReq *dtos.BookUpdateRequest) (*models.Book, error) {
	book, err := s.GetBookById(id)
	if err != nil {
		return nil, err
	}

	updatedBook := models.Book{
		ID:     bookReq.ID,
		Title:  bookReq.Title,
		Author: bookReq.Author,
		Year:   bookReq.Year,
		Publisher: bookReq.Publisher,
	}

	if bookReq.Cover != nil {
		cover := bookReq.Cover
		if book.Cover != "" {
			fmt.Println("Deleting old cover:", book.Cover)
			err := media.DeleteFileFromMinio(book.Cover)
			if err != nil {
				return nil, err
			}
		}

		fileReader, err := cover.Open()
		if err != nil {
			return nil, err
		}
		defer fileReader.Close()
		
		objectName, fileSize, contentType, err := media.ExtractFileData(cover)
		if err != nil {
			return nil, err
		}

		objectName = pathPrefix + objectName
		_, err = media.SendFileToMinio(objectName, fileReader, fileSize, contentType)
		if err != nil {
			return nil, err
		}

		updatedBook.Cover = objectName
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