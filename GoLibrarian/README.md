# GoLibrarian

1. To enhance understanding of Go programming language in the context of database integration.

2. To learn how to implement CRUD operations using PostgreSQL in Go.

3. To migrate from in-memory data storage to a persistent database.

4. To refine API documentation to include database integration aspects.



# Setting up PostgreSQL for Go Library System

This guide will help set up PostgreSQL for use with the Go Library System. PostgreSQL is a powerful open-source relational database management system that provides robust features and excellent performance.

## Installation

To install PostgreSQL on your system, use the following command:

```bash
sudo apt install postgresql postgresql-contrib
```

## Starting PostgreSQL Service

Once PostgreSQL is installed, you can start the PostgreSQL service using the following command:

```bash
sudo service postgresql start
```

## Checking Service Status

To check the status of the PostgreSQL service, use:

```bash
sudo service postgresql status
```

## Enabling Automatic Startup

To ensure that PostgreSQL starts automatically on system boot, enable it with:

```bash
sudo systemctl enable postgresql
```

```bash
createdb library_system
```

This command will create a new PostgreSQL database named `library_system`.

```bash
psql -U akhilesh -d library_system -h localhost
```

This command will connect to the `library_system` database using the user `akhilesh` and the host `localhost`.

```sql
-- Create a table to store information about books
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    published_date DATE,
    isbn TEXT UNIQUE
);
```

This SQL script will create a table named `books` with columns `id`, `title`, `author`, `published_date`, and `isbn`. The `id` column is of type SERIAL and serves as the primary key. The `title` and `author` columns are of type TEXT and cannot be NULL. The `published_date` column is of type DATE. The `isbn` column is of type TEXT with a UNIQUE constraint to ensure each ISBN is unique in the table.


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
   go get github.com/lib/pq

4. Run the Server:
   To run the GoLibraryAPI server, execute the following command:
   ```bash
   go run main.go

## Endpoint Descriptions:

Base URL: http://localhost:8080

The file contains a collection of HTTP requests designed to interact with a service that manages books:

1. **Get All Books**:
   - Method: GET
   - URL: http://localhost:8080/books

2. **Add Book**:
   - Method: POST
   - URL: http://localhost:8080/books
   - Body:
     ```json
     {
       "id": 1,
       "title": "Ikigai",
       "author": "Francesc Miralles",
       "published_date": "2016-01-01",
       "isbn": "453666"
     }
     ```

3. **Get Book by ID**:
   - Method: GET
   - URL: http://localhost:8080/books/1

4. **Update Book**:
   - Method: PUT
   - URL: http://localhost:8080/books/1
   - Body:
     ```json
     {
       "title": "Updated Ikigai",
       "author": "Akhilesh Sahu",
       "published_date": "2016-01-01",
       "isbn": "453667"
     }
     ```

5. **Delete Book**:
   - Method: DELETE
   - URL: http://localhost:8080/books/1

This configuration file facilitates CRUD (Create, Read, Update, Delete) operations on a collection of books through various HTTP requests. Each request is uniquely named and includes the method, URL, and, in the case of POST and PUT requests, a request body.

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
- **CRUD Operations:** Implemented Create, Read, Update, Delete operations efficiently using DB.
- **Error Handling:** Learned to manage errors effectively for a robust API.
- **Documentation:** Created comprehensive documentation for clear usage.
- **Testing:** Provided Postman collection for easy API testing.

### Experiences:
- **Hands-on Development:** Applied theoretical knowledge practically.
- **Collaborative Work:** Enhanced teamwork and problem-solving skills.
- **Project Management:** Developed planning and organization skills.
- **Continuous Learning:** Kept learning and improving throughout the project.

In conclusion, the provided JSON represents a configuration file for a REST API client, Thunder Client. This file outlines various HTTP requests for interacting with a service managing books. The requests include operations such as retrieving all books, adding a new book, fetching a book by its ID, updating a book, and deleting a book. Each request specifies the HTTP method, URL, and, in the case of POST and PUT requests, a JSON payload containing book details. This configuration file serves as a comprehensive setup for testing and validating endpoints in a RESTful API that deals with book management.
