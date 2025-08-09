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

func TestGetAllBooks(t *testing.T) {
	tests := []struct {
		name           string
		setupMocks     func(*mocks.UseCase)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "successful get all books",
			setupMocks: func(uc *mocks.UseCase) {
				books := []domain.Book{
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
				}
				uc.On("GetAllBooks", mock.Anything).Return(books, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"data": []interface{}{
					map[string]interface{}{
						"id":     "1",
						"title":  "Book 1",
						"author": "Author 1",
						"year":   float64(2021),
					},
					map[string]interface{}{
						"id":     "2",
						"title":  "Book 2",
						"author": "Author 2",
						"year":   float64(2022),
					},
					map[string]interface{}{
						"id":     "3",
						"title":  "Book 3",
						"author": "Author 3",
						"year":   float64(2023),
					},
				},
			},
		},
		{
			name: "empty books list",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("GetAllBooks", mock.Anything).Return([]domain.Book{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"data":    []interface{}{},
			},
		},
		{
			name: "usecase error",
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("GetAllBooks", mock.Anything).Return(nil, errors.New("database connection error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "database connection error",
			},
		},
		{
			name: "single book result",
			setupMocks: func(uc *mocks.UseCase) {
				books := []domain.Book{
					{
						ID:     "single-id",
						Title:  "Single Book",
						Author: "Single Author",
						Year:   2023,
					},
				}
				uc.On("GetAllBooks", mock.Anything).Return(books, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"data": []interface{}{
					map[string]interface{}{
						"id":     "single-id",
						"title":  "Single Book",
						"author": "Single Author",
						"year":   float64(2023),
					},
				},
			},
		},
		{
			name: "books with special characters",
			setupMocks: func(uc *mocks.UseCase) {
				books := []domain.Book{
					{
						ID:     "special-1",
						Title:  "Title with 特殊字符",
						Author: "Author with éàü",
						Year:   2024,
					},
				}
				uc.On("GetAllBooks", mock.Anything).Return(books, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"data": []interface{}{
					map[string]interface{}{
						"id":     "special-1",
						"title":  "Title with 特殊字符",
						"author": "Author with éàü",
						"year":   float64(2024),
					},
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
			app.Get("/books", handler.GetAllBooks)

			req := httptest.NewRequest(http.MethodGet, "/books", nil)

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