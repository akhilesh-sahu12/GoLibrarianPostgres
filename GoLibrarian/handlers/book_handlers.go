package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"akhilesh.sahu/GoLibraryAPI/models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

// LogRequest logs incoming requests.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the request
		log.Printf(
			"%s %s %s %s",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

func InitializeDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllBooks returns all books.
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, author, published_date, isbn FROM books")
	if err != nil {
		handleError(w, fmt.Errorf("failed to fetch books: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.ISBN)
		if err != nil {
			handleError(w, fmt.Errorf("failed to scan book row: %v", err), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	writeJSONResponse(w, books)
}

// AddBook adds a new book.
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		handleError(w, fmt.Errorf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	// Insert book into database
	result, err := db.Exec("INSERT INTO books (title, author, published_date, isbn) VALUES ($1, $2, $3, $4)", book.Title, book.Author, book.PublishedDate, book.ISBN)
	if err != nil {
		handleError(w, fmt.Errorf("failed to insert book: %v", err), http.StatusInternalServerError)
		return
	}

	// Get the inserted book ID
	id, err := result.LastInsertId()
	if err != nil {
		handleError(w, fmt.Errorf("failed to get last insert ID: %v", err), http.StatusInternalServerError)
		return
	}

	book.ID = int(id)

	writeJSONResponse(w, book)
}

// GetBookByID retrieves a book by ID.
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, fmt.Errorf("invalid book ID: %v", err), http.StatusBadRequest)
		return
	}

	var book models.Book
	err = db.QueryRow("SELECT title, author, published_date, isbn FROM books WHERE id = $1", id).Scan(&book.Title, &book.Author, &book.PublishedDate, &book.ISBN)
	if err == sql.ErrNoRows {
		handleError(w, fmt.Errorf("book not found"), http.StatusNotFound)
		return
	}
	if err != nil {
		handleError(w, fmt.Errorf("failed to fetch book: %v", err), http.StatusInternalServerError)
		return
	}

	book.ID = id

	writeJSONResponse(w, book)
}

// UpdateBook updates a book by ID.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, fmt.Errorf("invalid book ID: %v", err), http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		handleError(w, fmt.Errorf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE books SET title = $1, author = $2, published_date = $3, isbn = $4 WHERE id = $5", updatedBook.Title, updatedBook.Author, updatedBook.PublishedDate, updatedBook.ISBN, id)
	if err != nil {
		handleError(w, fmt.Errorf("failed to update book: %v", err), http.StatusInternalServerError)
		return
	}

	updatedBook.ID = id

	writeJSONResponse(w, updatedBook)
}

// DeleteBook deletes a book by ID.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, fmt.Errorf("invalid book ID: %v", err), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		handleError(w, fmt.Errorf("failed to delete book: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// handleError writes an error response.
func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("Error: %v", err)
	http.Error(w, err.Error(), statusCode)
}

// writeJSONResponse writes a JSON response.
func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// CloseDB closes the database connection.
func CloseDB() {
	db.Close()
}
