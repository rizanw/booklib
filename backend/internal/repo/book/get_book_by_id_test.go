package book

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	domain "booklib/internal/domain/book"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetBookByID(t *testing.T) {
	tests := []struct {
		name         string
		bookID       string
		setupMocks   func(mock sqlmock.Sqlmock)
		expectedBook *domain.Book
		expectedErr  string
	}{
		{
			name:   "successful get book by id",
			bookID: "test-id",
			setupMocks: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "author", "year", "created_at", "updated_at"}).
					AddRow("test-id", "Test Book", "Test Author", 2023, time.Now(), time.Now())
				mock.ExpectQuery(`SELECT \* FROM books WHERE id = \$1`).
					WithArgs("test-id").
					WillReturnRows(rows)
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
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM books WHERE id = \$1`).
					WithArgs("non-existent-id").
					WillReturnError(sql.ErrNoRows)
			},
			expectedBook: nil,
			expectedErr:  "",
		},
		{
			name:   "database error",
			bookID: "test-id",
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM books WHERE id = \$1`).
					WithArgs("test-id").
					WillReturnError(errors.New("database connection error"))
			},
			expectedBook: nil,
			expectedErr:  "database connection error",
		},
		{
			name:   "scan error",
			bookID: "test-id",
			setupMocks: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "author", "year", "created_at", "updated_at"}).
					AddRow("test-id", "Test Book", "Test Author", "invalid-year", time.Now(), time.Now())
				mock.ExpectQuery(`SELECT \* FROM books WHERE id = \$1`).
					WithArgs("test-id").
					WillReturnRows(rows)
			},
			expectedBook: nil,
			expectedErr:  "converting driver.Value type string",
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

			book, err := repo.GetBookByID(context.Background(), tt.bookID)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				assert.Nil(t, book)
			} else {
				assert.NoError(t, err)
				if tt.expectedBook == nil {
					assert.Nil(t, book)
				} else {
					assert.NotNil(t, book)
					assert.Equal(t, tt.expectedBook.ID, book.ID)
					assert.Equal(t, tt.expectedBook.Title, book.Title)
					assert.Equal(t, tt.expectedBook.Author, book.Author)
					assert.Equal(t, tt.expectedBook.Year, book.Year)
				}
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}