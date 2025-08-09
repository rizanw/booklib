package urlprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("creates new usecase", func(t *testing.T) {

		uc := New()

		assert.NotNil(t, uc)
		assert.Implements(t, (*UseCase)(nil), uc)
	})
}
