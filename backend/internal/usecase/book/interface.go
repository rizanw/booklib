package book

import (
	"context"

	domain "booklib/internal/domain/book"
)

//go:generate mockery --name=UseCase --output=./mocks
type UseCase interface {
	GetBook(ctx context.Context, id string) (*domain.Book, error)
	GetAllBooks(ctx context.Context) ([]domain.Book, error)
	AddBook(ctx context.Context, in AddBookInput) error
	UpdateBook(ctx context.Context, id string, in UpdateBookInput) error
	DeleteBook(ctx context.Context, id string) error
}
