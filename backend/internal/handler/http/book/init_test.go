package book

import (
	"testing"

	"booklib/internal/usecase/book/mocks"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("creates new handler with usecase", func(t *testing.T) {
		usecase := mocks.NewUseCase(t)
		
		handler := New(usecase)
		
		assert.NotNil(t, handler)
		assert.Equal(t, usecase, handler.usecase)
	})

	t.Run("handler maintains usecase reference", func(t *testing.T) {
		usecase := mocks.NewUseCase(t)
		
		handler := New(usecase)
		
		assert.NotNil(t, handler.usecase)
		assert.Equal(t, usecase, handler.usecase)
	})
}