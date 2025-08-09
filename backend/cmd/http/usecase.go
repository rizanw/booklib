package main

import (
	"booklib/internal/usecase/book"
	"booklib/internal/usecase/url-processor"
)

type UseCase struct {
	Book         book.UseCase
	UrlProcessor urlprocessor.UseCase
}

func newUseCase(repo *Repo) *UseCase {
	return &UseCase{
		Book:         book.New(repo.Book),
		UrlProcessor: urlprocessor.New(),
	}
}
