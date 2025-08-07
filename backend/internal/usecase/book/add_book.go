package book

import (
	"booklib/internal/domain/book"
	"context"
)

type AddBookInput struct {
	Title  string
	Author string
	Year   int
}

func (u usecase) AddBook(ctx context.Context, in AddBookInput) error {
	bk, err := book.NewBook(in.Title, in.Author, in.Year)
	if err != nil {
		return err
	}

	return u.repo.AddBook(ctx, bk)
}
