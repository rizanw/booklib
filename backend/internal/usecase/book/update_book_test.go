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

func TestUpdateBook(t *testing.T) {
	tests := []struct {
		name        string
		bookID      string
		input       UpdateBookInput
		setupMocks  func(*mocks.Repository)
		expectedErr string
	}{
		{
			name:   "successful update book",
			bookID: "test-id",
			input: UpdateBookInput{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(repo *mocks.Repository) {
				existingBook := &domain.Book{
					ID:     "test-id",
					Title:  "Old Book",
					Author: "Old Author",
					Year:   2020,
				}
				repo.On("GetBookByID", context.Background(), "test-id").Return(existingBook, nil)
				repo.On("UpdateBook", context.Background(), mock.MatchedBy(func(book *domain.Book) bool {
					return book.ID == "test-id" && book.Title == "Updated Book" && 
						   book.Author == "Updated Author" && book.Year == 2024
				})).Return(nil)
			},
			expectedErr: "",
		},
		{
			name:   "book not found for update",
			bookID: "non-existent-id",
			input: UpdateBookInput{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(repo *mocks.Repository) {
				repo.On("GetBookByID", context.Background(), "non-existent-id").Return(nil, errors.New("book not found"))
			},
			expectedErr: "book not found",
		},
		{
			name:   "get book error",
			bookID: "test-id",
			input: UpdateBookInput{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(repo *mocks.Repository) {
				repo.On("GetBookByID", context.Background(), "test-id").Return(nil, errors.New("repository error"))
			},
			expectedErr: "repository error",
		},
		{
			name:   "update book repository error",
			bookID: "test-id",
			input: UpdateBookInput{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(repo *mocks.Repository) {
				existingBook := &domain.Book{
					ID:     "test-id",
					Title:  "Old Book",
					Author: "Old Author",
					Year:   2020,
				}
				repo.On("GetBookByID", context.Background(), "test-id").Return(existingBook, nil)
				repo.On("UpdateBook", context.Background(), mock.Anything).Return(errors.New("update failed"))
			},
			expectedErr: "update failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			tt.setupMocks(repo)

			uc := New(repo)
			err := uc.UpdateBook(context.Background(), tt.bookID, tt.input)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}