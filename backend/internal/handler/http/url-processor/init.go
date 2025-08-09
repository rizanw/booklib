package urlprocessor

import (
	urlprocessor "booklib/internal/usecase/url-processor"
)

type Handler struct {
	usecase urlprocessor.UseCase
}

func New(usecase urlprocessor.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
