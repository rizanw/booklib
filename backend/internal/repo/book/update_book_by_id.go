package book

import (
	domain "booklib/internal/domain/book"
	"context"
)

func (r *repo) UpdateBook(ctx context.Context, book *domain.Book) error {
	query := `UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4`

	_, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.Year, book.ID)

	return err
}
