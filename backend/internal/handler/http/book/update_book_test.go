package book

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	usecaseBook "booklib/internal/usecase/book"
	"booklib/internal/usecase/book/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateBook(t *testing.T) {
	tests := []struct {
		name           string
		bookID         string
		requestBody    interface{}
		setupMocks     func(*mocks.UseCase)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "successful update book",
			bookID: "test-id",
			requestBody: UpdateBookRequest{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("UpdateBook", mock.Anything, "test-id", usecaseBook.UpdateBookInput{
					Title:  "Updated Book",
					Author: "Updated Author",
					Year:   2024,
				}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"status": "success",
			},
		},
		{
			name:           "invalid JSON body",
			bookID:         "test-id",
			requestBody:    `{"title": "Updated Book", "author": "Updated Author", "year": "invalid"}`,
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "Cannot parse JSON",
			},
		},
		{
			name:   "empty title validation",
			bookID: "test-id",
			requestBody: UpdateBookRequest{
				Title:  "",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "title cannot be empty",
			},
		},
		{
			name:   "empty author validation",
			bookID: "test-id",
			requestBody: UpdateBookRequest{
				Title:  "Updated Book",
				Author: "",
				Year:   2024,
			},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "author cannot be empty",
			},
		},
		{
			name:   "zero year validation",
			bookID: "test-id",
			requestBody: UpdateBookRequest{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   0,
			},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "year cannot be empty",
			},
		},
		{
			name:   "usecase error",
			bookID: "test-id",
			requestBody: UpdateBookRequest{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("UpdateBook", mock.Anything, "test-id", mock.Anything).Return(errors.New("book not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "book not found",
			},
		},
		{
			name:   "update with special characters",
			bookID: "special-id",
			requestBody: UpdateBookRequest{
				Title:  "Updated Title with 特殊字符",
				Author: "Updated Author with éàü",
				Year:   2024,
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("UpdateBook", mock.Anything, "special-id", usecaseBook.UpdateBookInput{
					Title:  "Updated Title with 特殊字符",
					Author: "Updated Author with éàü",
					Year:   2024,
				}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"status": "success",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			usecase := mocks.NewUseCase(t)
			tt.setupMocks(usecase)

			handler := New(usecase)
			app.Put("/books/:id", handler.UpdateBook)

			var body []byte
			var err error

			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.requestBody)
				assert.NoError(t, err)
			}

			url := "/books/" + tt.bookID
			req := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			for key, _ := range tt.expectedBody {
				_, exists := responseBody[key]
				assert.True(t, exists, "Expected key %s not found in response", key)
			}
		})
	}
}

func TestUpdateBookRequest_parseValidateRequest(t *testing.T) {
	tests := []struct {
		name        string
		request     UpdateBookRequest
		expected    usecaseBook.UpdateBookInput
		expectedErr string
	}{
		{
			name: "valid request",
			request: UpdateBookRequest{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			expected: usecaseBook.UpdateBookInput{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   2024,
			},
			expectedErr: "",
		},
		{
			name: "empty title",
			request: UpdateBookRequest{
				Title:  "",
				Author: "Updated Author",
				Year:   2024,
			},
			expected:    usecaseBook.UpdateBookInput{},
			expectedErr: "title cannot be empty",
		},
		{
			name: "empty author",
			request: UpdateBookRequest{
				Title:  "Updated Book",
				Author: "",
				Year:   2024,
			},
			expected:    usecaseBook.UpdateBookInput{},
			expectedErr: "author cannot be empty",
		},
		{
			name: "zero year",
			request: UpdateBookRequest{
				Title:  "Updated Book",
				Author: "Updated Author",
				Year:   0,
			},
			expected:    usecaseBook.UpdateBookInput{},
			expectedErr: "year cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.request.parseValidateRequest()

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				assert.Equal(t, usecaseBook.UpdateBookInput{}, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
