package book

import (
	"database/sql"
	"testing"
	"time"

	domain "booklib/internal/domain/book"

	"github.com/stretchr/testify/assert"
)

func TestBook_ToDomain(t *testing.T) {
	tests := []struct {
		name         string
		repoBook     Book
		expectedBook *domain.Book
	}{
		{
			name: "successful conversion with all fields",
			repoBook: Book{
				ID:        "test-id",
				Title:     "Test Book",
				Author:    "Test Author",
				Year:      2023,
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			},
			expectedBook: &domain.Book{
				ID:     "test-id",
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
		},
		{
			name: "conversion with null timestamps",
			repoBook: Book{
				ID:        "test-id-2",
				Title:     "Another Book",
				Author:    "Another Author",
				Year:      2022,
				CreatedAt: sql.NullTime{Valid: false},
				UpdatedAt: sql.NullTime{Valid: false},
			},
			expectedBook: &domain.Book{
				ID:     "test-id-2",
				Title:  "Another Book",
				Author: "Another Author",
				Year:   2022,
			},
		},
		{
			name: "conversion with empty strings",
			repoBook: Book{
				ID:        "",
				Title:     "",
				Author:    "",
				Year:      0,
				CreatedAt: sql.NullTime{Valid: false},
				UpdatedAt: sql.NullTime{Valid: false},
			},
			expectedBook: &domain.Book{
				ID:     "",
				Title:  "",
				Author: "",
				Year:   0,
			},
		},
		{
			name: "conversion with special characters",
			repoBook: Book{
				ID:        "special-id-123",
				Title:     "Title with 特殊字符 & symbols!",
				Author:    "Author with éàü accents",
				Year:      2024,
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			},
			expectedBook: &domain.Book{
				ID:     "special-id-123",
				Title:  "Title with 特殊字符 & symbols!",
				Author: "Author with éàü accents",
				Year:   2024,
			},
		},
		{
			name: "conversion with negative year",
			repoBook: Book{
				ID:        "ancient-book",
				Title:     "Ancient Text",
				Author:    "Ancient Author",
				Year:      -500,
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			},
			expectedBook: &domain.Book{
				ID:     "ancient-book",
				Title:  "Ancient Text",
				Author: "Ancient Author",
				Year:   -500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domainBook := tt.repoBook.ToDomain()

			assert.NotNil(t, domainBook)
			assert.Equal(t, tt.expectedBook.ID, domainBook.ID)
			assert.Equal(t, tt.expectedBook.Title, domainBook.Title)
			assert.Equal(t, tt.expectedBook.Author, domainBook.Author)
			assert.Equal(t, tt.expectedBook.Year, domainBook.Year)
		})
	}
}

func TestBook_ToDomain_Immutability(t *testing.T) {
	t.Run("modifying domain book should not affect repository book", func(t *testing.T) {
		repoBook := Book{
			ID:     "immutable-test",
			Title:  "Original Title",
			Author: "Original Author",
			Year:   2023,
		}

		domainBook := repoBook.ToDomain()
		originalTitle := repoBook.Title

		domainBook.Title = "Modified Title"

		assert.Equal(t, originalTitle, repoBook.Title)
		assert.NotEqual(t, repoBook.Title, domainBook.Title)
	})
}