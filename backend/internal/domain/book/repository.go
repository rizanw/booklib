package book

import "context"

type Repository interface {
	AddBook(ctx context.Context, book *Book) error
	GetAllBooks(ctx context.Context) ([]Book, error)
	GetBookByID(ctx context.Context, id string) (*Book, error)
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBook(ctx context.Context, id string) error
}
