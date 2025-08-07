package book

import (
	domain "booklib/internal/domain/book"
)

type usecase struct {
	repo domain.Repository
}

func New(repo domain.Repository) UseCase {
	return &usecase{
		repo: repo,
	}
}
