package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithError(t *testing.T) {
	tests := []struct {
		name          string
		statusCode    int
		message       string
		expectedError string
	}{
		{"BadRequest Error", http.StatusBadRequest, "Test error", "Test error"},
		{"NotFound Error", http.StatusNotFound, "Not found", "Not found"},
		{"InternalServerError", http.StatusInternalServerError, "Server error", "Server error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			WithError(rec, tt.statusCode, tt.message)

			res := rec.Result()
			if res.StatusCode != tt.statusCode {
				t.Errorf("expected status code %v; got %v", tt.statusCode, res.StatusCode)
			}

			var body map[string]string
			json.NewDecoder(res.Body).Decode(&body)
			if body["error"] != tt.expectedError {
				t.Errorf("expected error message %q; got %q", tt.expectedError, body["error"])
			}
		})
	}
}

func TestWithJSON(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		data       map[string]string
		expected   map[string]string
	}{
		{"Simple JSON Response", http.StatusOK, map[string]string{"key": "value"}, map[string]string{"key": "value"}},
		{"Empty JSON Response", http.StatusOK, map[string]string{}, map[string]string{}},
		{"Multiple Key-Value JSON Response", http.StatusOK, map[string]string{"key1": "value1", "key2": "value2"}, map[string]string{"key1": "value1", "key2": "value2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			WithJSON(rec, tt.statusCode, tt.data)

			res := rec.Result()
			if res.StatusCode != tt.statusCode {
				t.Errorf("expected status code %v; got %v", tt.statusCode, res.StatusCode)
			}

			var body map[string]string
			json.NewDecoder(res.Body).Decode(&body)
			for k, v := range tt.expected {
				if body[k] != v {
					t.Errorf("expected %s value %q; got %q", k, v, body[k])
				}
			}
		})
	}
}
