package book

import (
	"context"
	"errors"
	"testing"

	domain "booklib/internal/domain/book"
	"booklib/internal/domain/book/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetBook(t *testing.T) {
	tests := []struct {
		name         string
		bookID       string
		setupMocks   func(*mocks.Repository)
		expectedBook *domain.Book
		expectedErr  string
	}{
		{
			name:   "successful get book",
			bookID: "test-id",
			setupMocks: func(repo *mocks.Repository) {
				expectedBook := &domain.Book{
					ID:     "test-id",
					Title:  "Test Book",
					Author: "Test Author",
					Year:   2023,
				}
				repo.On("GetBookByID", context.Background(), "test-id").Return(expectedBook, nil)
			},
			expectedBook: &domain.Book{
				ID:     "test-id",
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			expectedErr: "",
		},
		{
			name:   "book not found",
			bookID: "non-existent-id",
			setupMocks: func(repo *mocks.Repository) {
				repo.On("GetBookByID", context.Background(), "non-existent-id").Return(nil, errors.New("book not found"))
			},
			expectedBook: nil,
			expectedErr:  "book not found",
		},
		{
			name:   "repository error",
			bookID: "test-id",
			setupMocks: func(repo *mocks.Repository) {
				repo.On("GetBookByID", context.Background(), "test-id").Return(nil, errors.New("repository error"))
			},
			expectedBook: nil,
			expectedErr:  "repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			tt.setupMocks(repo)

			uc := New(repo)
			book, err := uc.GetBook(context.Background(), tt.bookID)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				assert.Nil(t, book)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBook, book)
			}
		})
	}
}