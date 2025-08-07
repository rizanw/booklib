package book

import (
	domain "booklib/internal/domain/book"
	"context"
)

func (u usecase) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	return u.repo.GetAllBooks(ctx)
}
