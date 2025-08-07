package book

import (
	domain "booklib/internal/domain/book"
	"context"
)

func (r *repo) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	var (
		query  = `SELECT * FROM books`
		result []domain.Book
	)

	var books []Book
	if err := r.db.SelectContext(ctx, &books, query); err != nil {
		return result, err
	}

	for _, book := range books {
		result = append(result, *book.ToDomain())
	}

	return result, nil
}
