package book

import (
	domain "booklib/internal/domain/book"
	"context"
)

type UseCase interface {
	GetBook(ctx context.Context, id string) (*domain.Book, error)
	GetAllBooks(ctx context.Context) ([]domain.Book, error)
	AddBook(ctx context.Context, in AddBookInput) error
	UpdateBook(ctx context.Context, id string, in UpdateBookInput) error
	DeleteBook(ctx context.Context, id string) error
}
