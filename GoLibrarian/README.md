# GoLibrarian

GoLibraryAPI is a simple RESTful API for managing a library system, implemented in Go.

## Setup

1. Navigate to the project directory:
   ```bash
   cd GoLibrarian

2. Initialize Go module:
   ```bash
   go mod init akhilesh.sahu/GoLibraryAPI

3. Install dependencies:
   ```bash
   go get -u github.com/gorilla/mux

4. Run the Server:
   To run the GoLibraryAPI server, execute the following command:
   ```bash
   go run main.go

## Test the API

Use Postman to test the API endpoints. Here are the available endpoints:

## Endpoint Descriptions:

Base URL: http://localhost:8080

### Get All Books

- **Method:** GET
- **URL:** `/books`
- **Description:** Retrieves a list of all books in the library.

### Add Book

- **Method:** POST
- **URL:** `/books`
- **Description:** Adds a new book to the library.

### Get Book by ID

- **Method:** GET
- **URL:** `/books/{id}`
- **Description:** Retrieves details of a specific book by its ID.

### Update Book

- **Method:** PUT
- **URL:** `/books/{id}`
- **Description:** Updates details of a specific book by its ID.

### Delete Book

- **Method:** DELETE
- **URL:** `/books/{id}`
- **Description:** Deletes a specific book from the library by its ID.







## Request and Response Formats


This section provides detailed documentation for the expected request and response formats for each endpoint in the Library API.

## Get All Books

### Request Format

- **Method:** GET
- **URL:** `http://localhost:8080/books`

### Response Format

- **Description:** Retrieves a list of all books in the library.
- **Response Body Format:** Array of JSON objects representing books. Each book object has the following structure:
  ```json
  [
  {
    "id": "integer",
    "title": "string",
    "author": "string",
    "year": "integer"
  } 
  ]

## Add Book

### Request Format

- **Method:** POST
- **URL:** `http://localhost:8080/books`
- **Request Body Format:** JSON object representing the new book. The object should have the following structure:
  ```json
  { 
    "id": "integer",
    "title": "string",
    "author": "string",
    "year": "integer"
  }
  
### Response Format

- **Description:** Adds a new book to the library.
- **Response Body Format:** JSON object representing the newly added book. The object has the following structure:
  ```json
  {
    "id": "integer",
    "title": "string",
    "author": "string",
    "year": "integer"
  }

## Get Book by ID

### Request Format

- **Method:** GET
- **URL:** `http://localhost:8080/books/{id}`
  - Replace `{id}` with the ID of the book to retrieve.

### Response Format

- **Description:** Retrieves details of a specific book by its ID.
- **Response Body Format:** JSON object representing the book. The object has the following structure:
  ```json
  {
    "id": "integer",
    "title": "string",
    "author": "string",
    "year": "integer"
  }
## Update Book

### Request Format

- **Method:** PUT
- **URL:** `http://localhost:8080/books/{id}`
  - Replace `{id}` with the ID of the book to update.
- **Request Body Format:** JSON object representing the updated book. The object should have the following structure:
  ```json
  {
    "title": "string",
    "author": "string",
    "year": "integer"
  }
### Response Format

- **Description:** Updates details of a specific book by its ID.
- **Response Body Format:** JSON object representing the updated book. The object has the following structure:
  ```json
  {
    "id": "integer",
    "title": "string",
    "author": "string",
    "year": "integer"
  }
## Delete Book

### Request Format

- **Method:** DELETE
- **URL:** `http://localhost:8080/books/{id}`
  - Replace `{id}` with the ID of the book to delete.

### Response Format

- **Description:** Deletes a specific book from the library by its ID.
- **Response Body Format:** This endpoint does not return a response body. The HTTP response status indicates the outcome of the deletion operation.


## Status Codes

### 200 OK
- **Description:** The request was successful.
- **Usage:** Returned for successful GET requests or successful updates (POST, PUT, DELETE) where a response body is not necessary.

### 201 Created
- **Description:** The resource was successfully created.
- **Usage:** Returned when a new resource is created, such as when adding a new book.

### 400 Bad Request
- **Description:** The request could not be understood by the server due to malformed syntax or invalid data.
- **Usage:** Returned when the request body is invalid or missing required parameters.

### 404 Not Found
- **Description:** The requested resource could not be found on the server.
- **Usage:** Returned when attempting to retrieve or manipulate a resource that does not exist, such as when fetching a book by ID that is not present in the library.

### 500 Internal Server Error
- **Description:** The server encountered an unexpected condition that prevented it from fulfilling the request.
- **Usage:** Returned for unexpected errors on the server side, such as database connection issues or internal logic errors.

These status codes help to communicate the outcome of the request to the client, indicating whether the operation was successful, encountered an error, or failed due to invalid input or missing resources.
## Conclusion

### Learnings:
- **Go Fundamentals:** Mastered basics like syntax, data types, and control structures.
- **REST Principles:** Understood URL design, HTTP methods, and status codes.
- **CRUD Operations:** Implemented Create, Read, Update, Delete operations efficiently.
- **Error Handling:** Learned to manage errors effectively for a robust API.
- **Documentation:** Created comprehensive documentation for clear usage.
- **Testing:** Provided Postman collection for easy API testing.

### Experiences:
- **Hands-on Development:** Applied theoretical knowledge practically.
- **Collaborative Work:** Enhanced teamwork and problem-solving skills.
- **Project Management:** Developed planning and organization skills.
- **Continuous Learning:** Kept learning and improving throughout the project.

In summary, the project was a valuable learning experience, equipping us with essential skills and insights into Go programming and RESTful API development.
