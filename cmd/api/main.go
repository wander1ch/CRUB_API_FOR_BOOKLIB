package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/wander1ch/CRUB_API_FOR_BOOKLIB/internal/handlers"
	"github.com/wander1ch/CRUB_API_FOR_BOOKLIB/internal/models"
	"github.com/wander1ch/CRUB_API_FOR_BOOKLIB/internal/storage"

	_ "github.com/lib/pq"
)

type postgresStoreAdapter struct {
    *storage.PostgresStorage
}

func (a postgresStoreAdapter) CreateBook(book *models.Book) (int, error) {
    err := a.PostgresStorage.CreateBook(book)
    if err != nil {
        return 0, err
    }
    return book.ID, nil
}
func main() {
    connStr := "host=db user=damir dbname=booksdb sslmode=disable password=1234"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Cannot connect to DB:", err)
    }
    defer db.Close()

    if err = db.Ping(); err != nil {
        log.Fatal("Cannot ping DB:", err)
    }

    pgStore := storage.NewPostgresStorage(db)
    store := postgresStoreAdapter{pgStore}
    bookHandler := handlers.NewBookHandler(store)

    mux := http.NewServeMux()
    mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            bookHandler.ListBooks(w, r)
        case "POST":
            bookHandler.CreateBook(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    mux.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            bookHandler.GetBook(w, r)
        case "PUT":
            bookHandler.UpdateBook(w, r)
        case "DELETE":
            bookHandler.DeleteBook(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal("Server failed:", err)
    }
}