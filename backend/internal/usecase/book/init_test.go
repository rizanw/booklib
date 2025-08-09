package book

import (
	"testing"

	"booklib/internal/domain/book/mocks"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("creates new usecase with repository", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		
		uc := New(repo)
		
		assert.NotNil(t, uc)
		assert.Implements(t, (*UseCase)(nil), uc)
	})
}