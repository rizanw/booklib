package book

import (
	domain "booklib/internal/domain/book"
	"context"
	"database/sql"
	"errors"
)

func (r *repo) GetBookByID(ctx context.Context, id string) (*domain.Book, error) {
	var (
		query = `SELECT * FROM books WHERE id = $1`
		book  Book
	)

	if err := r.db.GetContext(ctx, &book, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return book.ToDomain(), nil
}
