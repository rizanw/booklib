package book

import (
	"context"
)

func (u usecase) DeleteBook(ctx context.Context, id string) error {
	return u.DeleteBook(ctx, id)
}
