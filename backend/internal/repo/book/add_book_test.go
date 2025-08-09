package book

import (
	"context"
	"errors"
	"testing"

	domain "booklib/internal/domain/book"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	tests := []struct {
		name        string
		book        *domain.Book
		setupMocks  func(mock sqlmock.Sqlmock)
		expectedErr string
	}{
		{
			name: "successful add book",
			book: &domain.Book{
				ID:     "test-id",
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO books \(id, title, author, year\) VALUES \(\$1, \$2, \$3, \$4\)`).
					WithArgs("test-id", "Test Book", "Test Author", 2023).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: "",
		},
		{
			name: "database error",
			book: &domain.Book{
				ID:     "test-id",
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO books \(id, title, author, year\) VALUES \(\$1, \$2, \$3, \$4\)`).
					WithArgs("test-id", "Test Book", "Test Author", 2023).
					WillReturnError(errors.New("database error"))
			},
			expectedErr: "database error",
		},
		{
			name: "constraint violation error",
			book: &domain.Book{
				ID:     "duplicate-id",
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO books \(id, title, author, year\) VALUES \(\$1, \$2, \$3, \$4\)`).
					WithArgs("duplicate-id", "Test Book", "Test Author", 2023).
					WillReturnError(errors.New("duplicate key value violates unique constraint"))
			},
			expectedErr: "duplicate key value violates unique constraint",
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

			err = repo.AddBook(context.Background(), tt.book)

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