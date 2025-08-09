package book

import (
	"testing"

	domain "booklib/internal/domain/book"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("creates new repository with database connection", func(t *testing.T) {
		db, _, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		repo := New(sqlxDB)

		assert.NotNil(t, repo)
		assert.Implements(t, (*domain.Repository)(nil), repo)
	})

	t.Run("repository maintains database connection", func(t *testing.T) {
		db, _, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := sqlx.NewDb(db, "sqlmock")
		repo := New(sqlxDB)

		assert.NotNil(t, repo)
		assert.Implements(t, (*domain.Repository)(nil), repo)
	})
}
