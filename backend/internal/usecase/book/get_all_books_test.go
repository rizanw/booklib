package book

import (
	"context"
	"errors"
	"testing"

	domain "booklib/internal/domain/book"
	"booklib/internal/domain/book/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBooks(t *testing.T) {
	tests := []struct {
		name          string
		setupMocks    func(*mocks.Repository)
		expectedBooks []domain.Book
		expectedErr   string
	}{
		{
			name: "successful get all books",
			setupMocks: func(repo *mocks.Repository) {
				books := []domain.Book{
					{
						ID:     "1",
						Title:  "Book 1",
						Author: "Author 1",
						Year:   2021,
					},
					{
						ID:     "2",
						Title:  "Book 2",
						Author: "Author 2",
						Year:   2022,
					},
				}
				repo.On("GetAllBooks", context.Background()).Return(books, nil)
			},
			expectedBooks: []domain.Book{
				{
					ID:     "1",
					Title:  "Book 1",
					Author: "Author 1",
					Year:   2021,
				},
				{
					ID:     "2",
					Title:  "Book 2",
					Author: "Author 2",
					Year:   2022,
				},
			},
			expectedErr: "",
		},
		{
			name: "empty result",
			setupMocks: func(repo *mocks.Repository) {
				repo.On("GetAllBooks", context.Background()).Return([]domain.Book{}, nil)
			},
			expectedBooks: []domain.Book{},
			expectedErr:   "",
		},
		{
			name: "repository error",
			setupMocks: func(repo *mocks.Repository) {
				repo.On("GetAllBooks", context.Background()).Return(nil, errors.New("repository error"))
			},
			expectedBooks: nil,
			expectedErr:   "repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			tt.setupMocks(repo)

			uc := New(repo)
			books, err := uc.GetAllBooks(context.Background())

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				assert.Nil(t, books)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBooks, books)
			}
		})
	}
}