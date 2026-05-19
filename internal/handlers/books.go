package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/wander1ch/CRUB_API_FOR_BOOKLIB/internal/models"
	"github.com/wander1ch/CRUB_API_FOR_BOOKLIB/internal/storage"
)


type BookHandler struct {
	store storage.BookStorer
}

func NewBookHandler(store storage.BookStorer) *BookHandler {
	return &BookHandler{store: store}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	id, err := h.store.CreateBook(&book)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}
	book.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/books/"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	book, err := h.store.GetBook(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if book == nil {
		http.NotFound(w,r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err :=  strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}


	if err := h.store.UpdateBook(id, &book); err != nil {
		http.Error(w, "Server error", http.StatusBadRequest)
		return	
	}
	book.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.store.DeleteBook(id); err != nil {
		http.Error(w, "Server error", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.store.ListBooks()	
	if err != nil {
		http.Error(w, "Server error", http.StatusBadRequest)
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}