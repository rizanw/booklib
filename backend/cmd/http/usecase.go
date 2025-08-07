package main

import "booklib/internal/usecase/book"

type UseCase struct {
	Book book.UseCase
}

func newUseCase(repo *Repo) *UseCase {
	return &UseCase{
		Book: book.New(repo.Book),
	}
}
