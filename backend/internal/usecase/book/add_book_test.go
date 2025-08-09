package book

import (
	"context"
	"errors"
	"testing"

	domain "booklib/internal/domain/book"
	"booklib/internal/domain/book/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddBook(t *testing.T) {
	tests := []struct {
		name        string
		input       AddBookInput
		setupMocks  func(*mocks.Repository)
		expectedErr string
	}{
		{
			name: "successful add book",
			input: AddBookInput{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks: func(repo *mocks.Repository) {
				repo.On("AddBook", mock.Anything, mock.MatchedBy(func(book *domain.Book) bool {
					return book.Title == "Test Book" && book.Author == "Test Author" && book.Year == 2023
				})).Return(nil)
			},
			expectedErr: "",
		},
		{
			name: "empty title",
			input: AddBookInput{
				Title:  "",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks:  func(repo *mocks.Repository) {},
			expectedErr: "title cannot be empty",
		},
		{
			name: "empty author",
			input: AddBookInput{
				Title:  "Test Book",
				Author: "",
				Year:   2023,
			},
			setupMocks:  func(repo *mocks.Repository) {},
			expectedErr: "author cannot be empty",
		},
		{
			name: "repository error",
			input: AddBookInput{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks: func(repo *mocks.Repository) {
				repo.On("AddBook", mock.Anything, mock.Anything).Return(errors.New("repository error"))
			},
			expectedErr: "repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			tt.setupMocks(repo)

			uc := New(repo)
			err := uc.AddBook(context.Background(), tt.input)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}