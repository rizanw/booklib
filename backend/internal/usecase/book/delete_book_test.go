package book

import (
	"context"
	"errors"
	"testing"

	"booklib/internal/domain/book/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDeleteBook(t *testing.T) {
	tests := []struct {
		name        string
		bookID      string
		setupMocks  func(*mocks.Repository)
		expectedErr string
	}{
		{
			name:   "successful delete book",
			bookID: "test-id",
			setupMocks: func(repo *mocks.Repository) {
				repo.On("DeleteBook", context.Background(), "test-id").Return(nil)
			},
			expectedErr: "",
		},
		{
			name:   "book not found for delete",
			bookID: "non-existent-id",
			setupMocks: func(repo *mocks.Repository) {
				repo.On("DeleteBook", context.Background(), "non-existent-id").Return(errors.New("book not found"))
			},
			expectedErr: "book not found",
		},
		{
			name:   "repository error during delete",
			bookID: "test-id",
			setupMocks: func(repo *mocks.Repository) {
				repo.On("DeleteBook", context.Background(), "test-id").Return(errors.New("repository error"))
			},
			expectedErr: "repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)
			tt.setupMocks(repo)

			uc := New(repo)
			err := uc.DeleteBook(context.Background(), tt.bookID)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}