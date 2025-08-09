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

func TestUpdateBook(t *testing.T) {
	tests := []struct {
		name        string
		book        *domain.Book
		setupMocks  func(mock sqlmock.Sqlmock)
		expectedErr string
	}{
		{
			name: "successful update book",
			book: &domain.Book{
				ID:     "test-id",
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE books SET title = \$1, author = \$2, year = \$3 WHERE id = \$4`).
					WithArgs("Updated Book", "Updated Author", 2024, "test-id").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: "",
		},
		{
			name: "no rows affected (book not found)",
			book: &domain.Book{
				ID:     "non-existent-id",
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE books SET title = \$1, author = \$2, year = \$3 WHERE id = \$4`).
					WithArgs("Updated Book", "Updated Author", 2024, "non-existent-id").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: "",
		},
		{
			name: "database error",
			book: &domain.Book{
				ID:     "test-id",
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE books SET title = \$1, author = \$2, year = \$3 WHERE id = \$4`).
					WithArgs("Updated Book", "Updated Author", 2024, "test-id").
					WillReturnError(errors.New("database connection error"))
			},
			expectedErr: "database connection error",
		},
		{
			name: "constraint violation error",
			book: &domain.Book{
				ID:     "test-id",
				Title:  "",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE books SET title = \$1, author = \$2, year = \$3 WHERE id = \$4`).
					WithArgs("", "Updated Author", 2024, "test-id").
					WillReturnError(errors.New("null value in column violates not-null constraint"))
			},
			expectedErr: "null value in column violates not-null constraint",
		},
		{
			name: "update with same values",
			book: &domain.Book{
				ID:     "test-id",
				Title:  "Same Title",
				Author: "Same Author",
				Year:   2023,
			},
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE books SET title = \$1, author = \$2, year = \$3 WHERE id = \$4`).
					WithArgs("Same Title", "Same Author", 2023, "test-id").
					WillReturnResult(sqlmock.NewResult(0, 1))
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

			err = repo.UpdateBook(context.Background(), tt.book)

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