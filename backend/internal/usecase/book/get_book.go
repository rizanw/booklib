package book

import (
	domain "booklib/internal/domain/book"
	"context"
)

func (u usecase) GetBook(ctx context.Context, id string) (*domain.Book, error) {
	return u.repo.GetBookByID(ctx, id)
}
