package storage

import (
	"database/sql"

	"github.com/wander1ch/CRUB_API_FOR_BOOKLIB/internal/models"
)
type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) CreateBook(book *models.Book) error {
	return s.db.QueryRow(
		"INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.Year,
	).Scan(&book.ID)
}

func (s *PostgresStorage) GetBook(id int) (*models.Book, error) {
	book := &models.Book{}
	err := s.db.QueryRow(
		"SELECT id, title, author, year FROM books WHERE id = $1",
		id,
	).Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *PostgresStorage) UpdateBook(id int, book *models.Book) error {
	_, err := s.db.Exec(
		"UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4",
		book.Title, book.Author, book.Year, id,
	)
	return err
}

func (s *PostgresStorage) DeleteBook(id int) error {
	_, err := s.db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}

func (s *PostgresStorage) ListBooks() ([]models.Book, error) {
	rows, err := s.db.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		book := models.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}