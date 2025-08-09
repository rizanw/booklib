package book

import (
	"context"
)

func (u usecase) DeleteBook(ctx context.Context, id string) error {
	return u.repo.DeleteBook(ctx, id)
}
