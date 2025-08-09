package urlprocessor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		operation   string
		expected    string
		expectedErr string
	}{
		{
			name:      "canonical operation - removes query and trailing slash",
			url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
			operation: "canonical",
			expected:  "https://BYFOOD.com/food-EXPeriences",
		},
		{
			name:      "redirection operation - converts to www.byfood.com and lowercase path",
			url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
			operation: "redirection",
			expected:  "https://www.byfood.com/food-experiences?query=abc/",
		},
		{
			name:      "all operation - applies both canonical and redirection",
			url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
			operation: "all",
			expected:  "https://www.byfood.com/food-experiences",
		},
		{
			name:      "canonical with different URL",
			url:       "http://example.com/path/to/resource/?param=value&other=123",
			operation: "canonical",
			expected:  "http://example.com/path/to/resource",
		},
		{
			name:      "redirection with different URL",
			url:       "http://example.com/PATH/TO/RESOURCE/?param=value",
			operation: "redirection",
			expected:  "https://www.byfood.com/path/to/resource/?param=value",
		},
		{
			name:      "all with different URL",
			url:       "http://example.com/PATH/TO/RESOURCE/?param=value&other=123",
			operation: "all",
			expected:  "https://www.byfood.com/path/to/resource",
		},
		{
			name:      "canonical with no trailing slash",
			url:       "https://test.com/path?query=value",
			operation: "canonical",
			expected:  "https://test.com/path",
		},
		{
			name:      "canonical with only trailing slash",
			url:       "https://test.com/?query=value",
			operation: "canonical",
			expected:  "https://test.com",
		},
		{
			name:      "redirection with uppercase path",
			url:       "https://test.com/UPPERCASE/PATH",
			operation: "redirection",
			expected:  "https://www.byfood.com/uppercase/path",
		},
		{
			name:        "invalid operation",
			url:         "https://BYFOOD.com/food-EXPeriences?query=abc/",
			operation:   "invalid",
			expectedErr: "invalid operation",
		},
		{
			name:        "invalid URL",
			url:         "://invalid-url",
			operation:   "all",
			expectedErr: "missing protocol scheme",
		},
		{
			name:      "empty path with canonical",
			url:       "https://example.com/?param=value",
			operation: "canonical",
			expected:  "https://example.com",
		},
		{
			name:      "empty path with redirection",
			url:       "https://example.com/?param=value",
			operation: "redirection",
			expected:  "https://www.byfood.com/?param=value",
		},
		{
			name:      "empty path with all",
			url:       "https://example.com/?param=value",
			operation: "all",
			expected:  "https://www.byfood.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New()
			result, err := uc.CleanURL(context.Background(), tt.operation, tt.url)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestCleanURLEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		operation   string
		expected    string
		expectedErr string
	}{
		{
			name:      "URL with fragment",
			url:       "https://example.com/path#fragment?query=value",
			operation: "canonical",
			expected:  "https://example.com/path#fragment?query=value",
		},
		{
			name:      "URL with port",
			url:       "https://example.com:8080/path?query=value",
			operation: "redirection",
			expected:  "https://www.byfood.com/path?query=value",
		},
		{
			name:      "URL with userinfo",
			url:       "https://user:pass@example.com/path?query=value",
			operation: "all",
			expected:  "https://user:pass@www.byfood.com/path",
		},
		{
			name:      "case insensitive operation",
			url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
			operation: "CANONICAL",
			expected:  "https://BYFOOD.com/food-EXPeriences",
		},
		{
			name:      "case insensitive operation - ALL",
			url:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
			operation: "ALL",
			expected:  "https://www.byfood.com/food-experiences",
		},
		{
			name:      "URL with multiple trailing slashes",
			url:       "https://example.com/path////?query=value",
			operation: "canonical",
			expected:  "https://example.com/path///",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New()
			result, err := uc.CleanURL(context.Background(), tt.operation, tt.url)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
