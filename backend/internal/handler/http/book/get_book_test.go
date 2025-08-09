package book

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	domain "booklib/internal/domain/book"
	"booklib/internal/usecase/book/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBook(t *testing.T) {
	tests := []struct {
		name           string
		bookID         string
		setupMocks     func(*mocks.UseCase)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "successful get book",
			bookID: "test-id",
			setupMocks: func(uc *mocks.UseCase) {
				book := &domain.Book{
					ID:     "test-id",
					Title:  "Test Book",
					Author: "Test Author",
					Year:   2023,
				}
				uc.On("GetBook", mock.Anything, "test-id").Return(book, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"data": map[string]interface{}{
					"id":     "test-id",
					"title":  "Test Book",
					"author": "Test Author",
					"year":   float64(2023),
				},
			},
		},
		{
			name:   "usecase error",
			bookID: "test-id",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("GetBook", mock.Anything, "test-id").Return(nil, errors.New("book not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "book not found",
			},
		},
		{
			name:   "database error",
			bookID: "test-id",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("GetBook", mock.Anything, "test-id").Return(nil, errors.New("database connection error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "database connection error",
			},
		},
		{
			name:   "book with special characters",
			bookID: "special-id-123",
			setupMocks: func(uc *mocks.UseCase) {
				book := &domain.Book{
					ID:     "special-id-123",
					Title:  "Title with 特殊字符 & symbols!",
					Author: "Author with éàü accents",
					Year:   2024,
				}
				uc.On("GetBook", mock.Anything, "special-id-123").Return(book, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"data": map[string]interface{}{
					"id":     "special-id-123",
					"title":  "Title with 特殊字符 & symbols!",
					"author": "Author with éàü accents",
					"year":   float64(2024),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			usecase := mocks.NewUseCase(t)
			tt.setupMocks(usecase)

			handler := New(usecase)
			app.Get("/books/:id", handler.GetBook)

			url := "/books/" + tt.bookID
			req := httptest.NewRequest(http.MethodGet, url, nil)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			for key, expectedValue := range tt.expectedBody {
				actualValue, exists := responseBody[key]
				assert.True(t, exists, "Expected key %s not found in response", key)
				assert.Equal(t, expectedValue, actualValue)
			}
		})
	}
}
