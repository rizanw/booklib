package urlprocessor

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"booklib/internal/usecase/url-processor/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProcessUrl(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupMocks     func(*mocks.UseCase)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "successful process URL with all operation",
			requestBody: ProcessUrlRequest{
				Url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "all",
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("CleanURL", mock.Anything, "all", "https://BYFOOD.com/food-EXPeriences?query=abc/").
					Return("https://www.byfood.com/food-experiences", nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"processed_url": "https://www.byfood.com/food-experiences",
			},
		},
		{
			name: "successful process URL with canonical operation",
			requestBody: ProcessUrlRequest{
				Url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "canonical",
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("CleanURL", mock.Anything, "canonical", "https://BYFOOD.com/food-EXPeriences?query=abc/").
					Return("https://BYFOOD.com/food-EXPeriences", nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"processed_url": "https://BYFOOD.com/food-EXPeriences",
			},
		},
		{
			name: "successful process URL with redirection operation",
			requestBody: ProcessUrlRequest{
				Url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "redirection",
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("CleanURL", mock.Anything, "redirection", "https://BYFOOD.com/food-EXPeriences?query=abc/").
					Return("https://www.byfood.com/food-experiences?query=abc/", nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"processed_url": "https://www.byfood.com/food-experiences?query=abc/",
			},
		},
		{
			name:           "invalid JSON body",
			requestBody:    `{"url": "https://example.com", "operation": 123}`,
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Cannot parse JSON",
			},
		},
		{
			name: "empty URL validation",
			requestBody: ProcessUrlRequest{
				Url:       "",
				Operation: "all",
			},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": map[string]interface{}{},
			},
		},
		{
			name: "empty operation validation",
			requestBody: ProcessUrlRequest{
				Url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "",
			},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": map[string]interface{}{},
			},
		},
		{
			name: "invalid operation from usecase",
			requestBody: ProcessUrlRequest{
				Url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "invalid",
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("CleanURL", mock.Anything, "invalid", "https://BYFOOD.com/food-EXPeriences?query=abc/").
					Return("", errors.New("invalid operation"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error": "invalid operation",
			},
		},
		{
			name: "invalid URL from usecase",
			requestBody: ProcessUrlRequest{
				Url:       "://invalid-url",
				Operation: "all",
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("CleanURL", mock.Anything, "all", "://invalid-url").
					Return("", errors.New("parse \"://invalid-url\": missing protocol scheme"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error": "parse \"://invalid-url\": missing protocol scheme",
			},
		},
		{
			name: "usecase internal error",
			requestBody: ProcessUrlRequest{
				Url:       "https://example.com",
				Operation: "all",
			},
			setupMocks: func(uc *mocks.UseCase) {
				uc.On("CleanURL", mock.Anything, "all", "https://example.com").
					Return("", errors.New("internal server error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error": "internal server error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			usecase := mocks.NewUseCase(t)
			tt.setupMocks(usecase)

			handler := New(usecase)
			app.Post("/process-url", handler.ProcessUrl)

			var body []byte
			var err error

			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.requestBody)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/process-url", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

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

func TestProcessUrlRequest_validate(t *testing.T) {
	tests := []struct {
		name        string
		request     ProcessUrlRequest
		expectedErr string
	}{
		{
			name: "valid request",
			request: ProcessUrlRequest{
				Url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "all",
			},
			expectedErr: "",
		},
		{
			name: "empty URL",
			request: ProcessUrlRequest{
				Url:       "",
				Operation: "all",
			},
			expectedErr: "url cannot be empty",
		},
		{
			name: "empty operation",
			request: ProcessUrlRequest{
				Url:       "https://example.com",
				Operation: "",
			},
			expectedErr: "operation cannot be empty",
		},
		{
			name: "both fields empty",
			request: ProcessUrlRequest{
				Url:       "",
				Operation: "",
			},
			expectedErr: "url cannot be empty",
		},
		{
			name: "valid canonical operation",
			request: ProcessUrlRequest{
				Url:       "https://test.com/path?query=value",
				Operation: "canonical",
			},
			expectedErr: "",
		},
		{
			name: "valid redirection operation",
			request: ProcessUrlRequest{
				Url:       "https://test.com/path",
				Operation: "redirection",
			},
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.validate()

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProcessUrlEdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupMocks     func(*mocks.UseCase)
		expectedStatus int
		checkResponse  func(t *testing.T, responseBody map[string]interface{})
	}{
		{
			name: "malformed JSON",
			requestBody: `{
				"url": "https://example.com",
				"operation": "all"
				// missing closing brace and has comment
			`,
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, responseBody map[string]interface{}) {
				assert.Contains(t, responseBody, "error")
				assert.Equal(t, "Cannot parse JSON", responseBody["error"])
			},
		},
		{
			name: "empty request body",
			requestBody: ProcessUrlRequest{},
			setupMocks:     func(uc *mocks.UseCase) {},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, responseBody map[string]interface{}) {
				assert.Contains(t, responseBody, "error")
			},
		},
		{
			name: "very long URL",
			requestBody: ProcessUrlRequest{
				Url:       "https://example.com/" + string(make([]byte, 2000)),
				Operation: "all",
			},
			setupMocks: func(uc *mocks.UseCase) {
				longUrl := "https://example.com/" + string(make([]byte, 2000))
				uc.On("CleanURL", mock.Anything, "all", longUrl).
					Return("https://www.byfood.com/"+string(make([]byte, 2000)), nil)
			},
			expectedStatus: http.StatusCreated,
			checkResponse: func(t *testing.T, responseBody map[string]interface{}) {
				assert.Contains(t, responseBody, "processed_url")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			usecase := mocks.NewUseCase(t)
			tt.setupMocks(usecase)

			handler := New(usecase)
			app.Post("/process-url", handler.ProcessUrl)

			var body []byte
			var err error

			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.requestBody)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/process-url", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)

			tt.checkResponse(t, responseBody)
		})
	}
}