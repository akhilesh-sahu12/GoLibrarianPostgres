package main

import (
	"log"
	"net/http"
	"os"

	"akhilesh.sahu/GoLibraryAPI/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Create a connection string
	psqlInfo := "host=localhost port=5432 user=akhilesh " +
		"password=root dbname=library_system sslmode=disable"

	handlers.InitializeDB(psqlInfo)
	defer handlers.CloseDB()

	// Open a connection to the database
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatalf("Error connecting to the database: %v", err)
	// }
	// defer db.Close()

	// // Verify the connection
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("Error pinging the database: %v", err)
	// }

	// log.Println("Successfully connected to the database!")

	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books", handlers.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Log requests to stdout
	loggedRouter := handlers.LogRequest(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Printf("Server started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, loggedRouter))
}
