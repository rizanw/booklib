package book

import (
	domain "booklib/internal/domain/book"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) domain.Repository {
	return &repo{
		db: db,
	}
}
