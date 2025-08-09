package book

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "booklib/internal/domain/book"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBooks(t *testing.T) {
	tests := []struct {
		name          string
		setupMocks    func(mock sqlmock.Sqlmock)
		expectedBooks []domain.Book
		expectedErr   string
	}{
		{
			name: "successful get all books",
			setupMocks: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "author", "year", "created_at", "updated_at"}).
					AddRow("1", "Book 1", "Author 1", 2021, time.Now(), time.Now()).
					AddRow("2", "Book 2", "Author 2", 2022, time.Now(), time.Now()).
					AddRow("3", "Book 3", "Author 3", 2023, time.Now(), time.Now())
				mock.ExpectQuery(`SELECT \* FROM books`).
					WillReturnRows(rows)
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
				{
					ID:     "3",
					Title:  "Book 3",
					Author: "Author 3",
					Year:   2023,
				},
			},
			expectedErr: "",
		},
		{
			name: "empty result",
			setupMocks: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "author", "year", "created_at", "updated_at"})
				mock.ExpectQuery(`SELECT \* FROM books`).
					WillReturnRows(rows)
			},
			expectedBooks: []domain.Book{},
			expectedErr:   "",
		},
		{
			name: "database error",
			setupMocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM books`).
					WillReturnError(errors.New("database connection error"))
			},
			expectedBooks: []domain.Book{},
			expectedErr:   "database connection error",
		},
		{
			name: "scan error",
			setupMocks: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "author", "year", "created_at", "updated_at"}).
					AddRow("1", "Book 1", "Author 1", "invalid-year", time.Now(), time.Now())
				mock.ExpectQuery(`SELECT \* FROM books`).
					WillReturnRows(rows)
			},
			expectedBooks: []domain.Book{},
			expectedErr:   "converting driver.Value type string",
		},
		{
			name: "single book result",
			setupMocks: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "author", "year", "created_at", "updated_at"}).
					AddRow("single-id", "Single Book", "Single Author", 2023, time.Now(), time.Now())
				mock.ExpectQuery(`SELECT \* FROM books`).
					WillReturnRows(rows)
			},
			expectedBooks: []domain.Book{
				{
					ID:     "single-id",
					Title:  "Single Book",
					Author: "Single Author",
					Year:   2023,
				},
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

			books, err := repo.GetAllBooks(context.Background())

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedBooks), len(books))
				for i, expectedBook := range tt.expectedBooks {
					assert.Equal(t, expectedBook.ID, books[i].ID)
					assert.Equal(t, expectedBook.Title, books[i].Title)
					assert.Equal(t, expectedBook.Author, books[i].Author)
					assert.Equal(t, expectedBook.Year, books[i].Year)
				}
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}