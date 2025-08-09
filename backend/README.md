# 📦 Backend - Book Library CRUD App

This is a golang backend service for managing a small library of books and URL Cleanup and Redirection Service.

---

## 🛠️ Setup

### 1. Prerequisites

- Go ≥ 1.24
- postgresql (you can use docker-compose inside)

### 1.1 Database Prerequisites

Do database migration using scripts inside `migrations` folder.

```bash 
./create-db.sh 
./migrate-up.sh 
```

or you can migrate the `20250807_01_create_books_table.sql` and other sql files inside `migrations/up` folder manually.

### 2. Environment Variables

Update `config.yaml` file inside `files/etc/booklib`

### 3. Run Backend Server

```bash
go mod tidy
go run cmd/http/main.go
```

or using make:

```bash
make run
```

## 🗂️ Project Structure

```
⏺ backend/
  ├── bin/                    # Binary/executable files
  ├── cmd/                    # Application entry points
  │   └── http/              # HTTP server entry point
  ├── docs/                  # API documentation (Swagger)
  ├── files/                 # Configuration and static files
  │   └── etc/
  │       └── booklib/
  ├── internal/              # Private application code
  │   ├── domain/            # Business entities and interfaces
  │   │   └── book/
  │   │       └── mocks/     # Mock implementations
  │   ├── handler/           # HTTP request handlers
  │   │   └── http/
  │   │       ├── book/      # Book-related endpoints
  │   │       └── url-processor/  # URL processing endpoints
  │   ├── infra/             # Infrastructure layer
  │   │   └── config/        # Configuration management
  │   ├── repo/              # Data repository layer
  │   │   └── book/          # Book data operations
  │   └── usecase/           # Business logic layer
  │       ├── book/          # Book business logic
  │       │   └── mocks/     # Mock implementations
  │       └── url-processor/ # URL processing logic
  │           └── mocks/     # Mock implementations
  └── migrations/            # Database migration scripts
      ├── down/              # Rollback migrations
      └── up/                # Forward migrations
```

## 📄 API Documentation

The API has two main parts: Books CRUD and URL Cleanup & Redirection Service.

For full interactive API documentation, after starting the application, open:

```
http://localhost:<PORT>/docs
```

> This Swagger UI is auto-generated from code annotations.

### ✴ Books CRUD API

#### GET /api/v1/books

Retrieve all books.

**Response:**

```json
{
  "data": [
    {
      "id": "fbb7f0dd-2982-4023-b95e-0b97e09f53ce",
      "title": "Robert C. Martin",
      "author": "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
      "year": 2017
    }
  ],
  "status": "success"
}
```

#### GET /api/v1/books/{id}

Retrieve a single book by ID.

**Response:**

```json
{
  "data": {
    "id": "fbb7f0dd-2982-4023-b95e-0b97e09f53ce",
    "title": "Robert C. Martin",
    "author": "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
    "year": 2017
  },
  "status": "success"
}
```

#### POST /api/v1/books

Add a new book.

**Request:**

```json
{
  "title": "Robert C. Martin",
  "author": "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
  "year": 2017
}
```

**Response:**

```json
{
  "status": "success"
}
```

#### PUT /api/v1/books/{id}

Update a book by ID.

**Request:**

```json
{
  "title": "Robert C. Martin",
  "author": "Clean Architecture: A Craftsman's Guide to Software Structure and Design",
  "year": 2017
}
```

**Response:**

```json
{
  "status": "success"
}
```

#### DELETE /api/v1/books/{id}

Delete a book by ID.

**Response:**

```json
{
  "status": "success"
}
```

### ✴ URL Cleanup & Redirection Service API

#### POST /process-url

Operations:

- canonical → Remove query params & trailing slashes, keep host as-is
- redirection → Force domain to "www.byfood.com" and lowercase entire URL
- all → Apply canonical first, then redirection

**Request:**

```json
{
  "url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
  "operation": "all"
}
```

**Response:**

```json
{
  "processed_url": "https://www.byfood.com/food-experiences"
}
```

## 🧪 Testing

### Mocks

Most of the mocks are generated using [mockery](https://vektra.github.io/mockery/latest/). You can you comment command
just like this:

```go
//go:generate mockery --name=Repository --output=./mocks
type Repository interface {
AddBook(ctx context.Context, book *Book) error
GetAllBooks(ctx context.Context) ([]Book, error)
GetBookByID(ctx context.Context, id string) (*Book, error)
UpdateBook(ctx context.Context, book *Book) error
DeleteBook(ctx context.Context, id string) error
}
```

### Run Test

```bash
go test ./... -v 
```

or using make:

```bash
make test
```

### Test Coverage

to generate the test coverage, you can run:

```bash
go test ./... -coverprofile=coverage.out -v
```

or using make:

```bash
make coverage
```

or

```bash
make coveragetext
```
