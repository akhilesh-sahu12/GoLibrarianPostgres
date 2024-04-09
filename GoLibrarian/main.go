package main

import (
	"akhilesh.sahu/GoLibraryAPI/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books", handlers.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Log requests
	loggedRouter := handlers.LogRequest(router)

	// Start server
	log.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
