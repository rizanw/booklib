package book

import (
	domain "booklib/internal/domain/book"
	"context"
)

func (r *repo) AddBook(ctx context.Context, book *domain.Book) error {
	query := `INSERT INTO books (id, title, author, year) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, book.ID, book.Title, book.Author, book.Year)

	return err
}
