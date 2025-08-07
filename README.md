# Fullstack Book Library CRUD App

A full-fledged CRUD application that manages a small library of books.

## ✨ Features

### 📖 Book Library

- List all books
- Add a new book
- Edit book details
- View book details
- Delete a book
- Client-side form validation
- Modal-based forms
- Friendly error handling

## 🔧 Tech Stack

### Frontend:

- TypeScript

### Backend:

- Golang
- RESTful API
- PostgreSQL
- Contextual Logging with `github.com/rizanw/go-log` package
- Swagger/OpenAPI for API documentation

## 🧱 Domain-Driven Design (DDD) Overview

The backend is designed using **Clean Architecture / DDD principles**.

### 📘 Domain: Book

**Entity: `Book`**

```go
type Book struct {
ID     string // UUID
Title  string
Author string
Year   int
}

```

**Business Rules**:

- Title and Author must not be empty.

## 🧰 Setup Instructions

To set up and run the project locally, please refer to the setup instructions inside each subdirectory:

- 📦 Backend: [backend/README.md](backend/README.md)

- 💻 Frontend: [frontend/README.md](frontend/README.md)