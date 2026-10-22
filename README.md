# Book CRUD API (Go + Fiber)

A simple REST API for managing books, built with [Go](https://go.dev/) and the [Fiber v2](https://github.com/gofiber/fiber) web framework. Data is stored **in memory**, so it resets every time the server restarts.

## Project Structure

```
.
├── main.go              # App setup, Book model, seed data, routes
└── func-book-CRUD.go    # Handlers for the CRUD operations
```

| File | Description |
|------|-------------|
| `main.go` | Defines the `Book` struct, initializes sample data, registers routes, and starts the server on port `8080`. |
| `func-book-CRUD.go` | Contains the handler functions: `getBooks`, `getBook`, `createBook`, `updateBook`, `deleteBook`. |

## Requirements

- Go 1.18 or later
- Fiber v2

## Getting Started

```bash
# 1. Clone the repository
git clone <your-repo-url>
cd <your-repo-folder>

# 2. Initialize the module (if there is no go.mod yet)
go mod init book-api

# 3. Install dependencies
go mod tidy

# 4. Run the server
go run .
```

The server starts at `http://localhost:8080`.

> Note: use `go run .` (not `go run main.go`) so both files are compiled together.

## Data Model

```go
type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}
```

### Initial Data

| ID | Title | Author |
|----|-------|--------|
| 1 | 1984 | George Orwell |
| 2 | The Great Gatsby | F. Scott Fitzgerald |
| 3 | The Forest | Steam |

## API Endpoints

| Method | Endpoint | Description | Success Response |
|--------|----------|-------------|------------------|
| GET | `/books` | Get all books | `200 OK` + JSON array |
| GET | `/books/:id` | Get a book by ID | `200 OK` + JSON object |
| POST | `/books` | Create a new book | `200 OK` + created book |
| PUT | `/books/:id` | Update a book's title and author | `200 OK` + updated book |
| DELETE | `/books/:id` | Delete a book | `204 No Content` |

### Error Responses

| Status | When |
|--------|------|
| `400 Bad Request` | The `id` is not a number, or the request body is invalid |
| `404 Not Found` | No book exists with the given ID |

## Usage Examples (curl)

**Get all books**
```bash
curl http://localhost:8080/books
```

**Get one book**
```bash
curl http://localhost:8080/books/1
```

**Create a book**
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"id": 4, "title": "Brave New World", "author": "Aldous Huxley"}'
```

**Update a book**
```bash
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Nineteen Eighty-Four", "author": "George Orwell"}'
```

**Delete a book**
```bash
curl -X DELETE http://localhost:8080/books/1
```

## Known Limitations

- Data is not persisted; it is lost when the server stops.
- `createBook` does not auto-generate IDs or check for duplicates, so you must supply a unique `id`.
- No input validation (empty title/author are accepted).
- The in-memory slice is not protected against concurrent access (no mutex).

## Possible Improvements

- Auto-increment IDs
- Add a database (SQLite, PostgreSQL, etc.)
- Add validation and a mutex for thread safety
- Add unit tests

## License

This project is for learning purposes.
