package book

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"booklib/internal/usecase/book/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteBook(t *testing.T) {
	tests := []struct {
		name           string
		bookID         string
		setupMocks     func(*mocks.UseCase)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "successful delete book",
			bookID: "test-id",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("DeleteBook", mock.Anything, "test-id").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
			},
		},
		{
			name:   "usecase error - book not found",
			bookID: "non-existent-id",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("DeleteBook", mock.Anything, "non-existent-id").Return(errors.New("book not found"))
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
				uc.On("DeleteBook", mock.Anything, "test-id").Return(errors.New("database connection error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "database connection error",
			},
		},
		{
			name:   "constraint violation error",
			bookID: "referenced-id",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("DeleteBook", mock.Anything, "referenced-id").Return(errors.New("foreign key constraint fails"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "foreign key constraint fails",
			},
		},
		{
			name:   "delete book with special id",
			bookID: "special-id-123",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("DeleteBook", mock.Anything, "special-id-123").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
			},
		},
		{
			name:   "delete book with uuid format",
			bookID: "550e8400-e29b-41d4-a716-446655440000",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("DeleteBook", mock.Anything, "550e8400-e29b-41d4-a716-446655440000").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			usecase := mocks.NewUseCase(t)
			tt.setupMocks(usecase)

			handler := New(usecase)
			app.Delete("/books/:id", handler.DeleteBook)

			url := "/books/" + tt.bookID
			req := httptest.NewRequest(http.MethodDelete, url, nil)

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
