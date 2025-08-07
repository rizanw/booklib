package main

type Handler struct{}

func newHandler(uc *UseCase) *Handler {
	return &Handler{}
}
