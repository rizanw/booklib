package book

import "booklib/internal/usecase/book"

type Handler struct {
	usecase book.UseCase
}

func New(usecase book.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
