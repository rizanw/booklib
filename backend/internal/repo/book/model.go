package book

import (
	domain "booklib/internal/domain/book"
	"database/sql"
)

type Book struct {
	ID        string       `db:"id"`
	Title     string       `db:"title"`
	Author    string       `db:"author"`
	Year      int          `db:"year"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (b *Book) ToDomain() *domain.Book {
	return &domain.Book{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}
}
