package book

import (
	"context"
)

type UpdateBookInput struct {
	Title  string
	Author string
	Year   int
}

func (u usecase) UpdateBook(ctx context.Context, id string, in UpdateBookInput) error {
	bk, err := u.GetBook(ctx, id)
	if err != nil {
		return err
	}

	bk.Title = in.Title
	bk.Author = in.Author
	bk.Year = in.Year

	return u.repo.UpdateBook(ctx, bk)
}
