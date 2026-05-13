package storage

import (
	"github.com/wander1ch/CRUB_API_FOR_BOOKLIB/internal/models"
)

type BookStorer interface {
	CreateBook(book *models.Book) (int, error)
	GetBook(id int) (*models.Book, error)
	UpdateBook(id int, book *models.Book) error
	DeleteBook(id int) error
	ListBooks() ([]models.Book, error)
}