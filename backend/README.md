# 📦 Backend - Book Library CRUD App

## 🛠️ Setup

### 1. Prerequisites

- Go ≥ 1.24
- postgresql (you can use docker-compose here)

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

## 📄 API Documentation

TBD

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
