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

func TestAddBook(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupMocks     func(*mocks.UseCase)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "successful add book",
			requestBody: AddBookRequest{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("AddBook", mock.Anything, usecaseBook.AddBookInput{
					Title:  "Test Book",
					Author: "Test Author",
					Year:   2023,
				}).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"status": "success",
			},
		},
		{
			name:           "invalid JSON body",
			requestBody:    `{"title": "Test Book", "author": "Test Author", "year": "invalid"}`,
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Cannot parse JSON",
			},
		},
		{
			name: "empty title validation",
			requestBody: AddBookRequest{
				Title:  "",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "title cannot be empty",
			},
		},
		{
			name: "empty author validation",
			requestBody: AddBookRequest{
				Title:  "Test Book",
				Author: "",
				Year:   2023,
			},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "author cannot be empty",
			},
		},
		{
			name: "zero year validation",
			requestBody: AddBookRequest{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   0,
			},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "year cannot be empty",
			},
		},
		{
			name: "usecase error",
			requestBody: AddBookRequest{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("AddBook", mock.Anything, mock.Anything).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"status": "error",
				"error":  "database error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			usecase := mocks.NewUseCase(t)
			tt.setupMocks(usecase)

			handler := New(usecase)
			app.Post("/books", handler.AddBook)

			var body []byte
			var err error

			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.requestBody)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
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

func TestAddBookRequest_parseValidateRequest(t *testing.T) {
	tests := []struct {
		name        string
		request     AddBookRequest
		expected    usecaseBook.AddBookInput
		expectedErr string
	}{
		{
			name: "valid request",
			request: AddBookRequest{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			expected: usecaseBook.AddBookInput{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2023,
			},
			expectedErr: "",
		},
		{
			name: "empty title",
			request: AddBookRequest{
				Title:  "",
				Author: "Test Author",
				Year:   2023,
			},
			expected:    usecaseBook.AddBookInput{},
			expectedErr: "title cannot be empty",
		},
		{
			name: "empty author",
			request: AddBookRequest{
				Title:  "Test Book",
				Author: "",
				Year:   2023,
			},
			expected:    usecaseBook.AddBookInput{},
			expectedErr: "author cannot be empty",
		},
		{
			name: "zero year",
			request: AddBookRequest{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   0,
			},
			expected:    usecaseBook.AddBookInput{},
			expectedErr: "year cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.request.parseValidateRequest()

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				assert.Equal(t, usecaseBook.AddBookInput{}, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
