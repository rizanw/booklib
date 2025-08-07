package book

import (
	"errors"
	"github.com/google/uuid"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Year   int
}

func NewBook(title, author string, year int) (*Book, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if author == "" {
		return nil, errors.New("author cannot be empty")
	}

	return &Book{
		ID:     uuid.NewString(),
		Title:  title,
		Author: author,
		Year:   year,
	}, nil
}
