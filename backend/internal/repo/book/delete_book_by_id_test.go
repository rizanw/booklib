package book

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBook(t *testing.T) {
	tests := []struct {
		name        string
		bookID      string
		setupMocks  func(mock sqlmock.Sqlmock)
		expectedErr string
	}{
		{
			name:   "successful delete book",
			bookID: "test-id",
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM books WHERE id = \$1`).
					WithArgs("test-id").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: "",
		},
		{
			name:   "no rows affected (book not found)",
			bookID: "non-existent-id",
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM books WHERE id = \$1`).
					WithArgs("non-existent-id").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: "",
		},
		{
			name:   "database error",
			bookID: "test-id",
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM books WHERE id = \$1`).
					WithArgs("test-id").
					WillReturnError(errors.New("database connection error"))
			},
			expectedErr: "database connection error",
		},
		{
			name:   "constraint violation error",
			bookID: "referenced-id",
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM books WHERE id = \$1`).
					WithArgs("referenced-id").
					WillReturnError(errors.New("foreign key constraint fails"))
			},
			expectedErr: "foreign key constraint fails",
		},
		{
			name:   "empty book id",
			bookID: "",
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM books WHERE id = \$1`).
					WithArgs("").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			repo := New(sqlxDB)

			tt.setupMocks(mock)

			err = repo.DeleteBook(context.Background(), tt.bookID)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}