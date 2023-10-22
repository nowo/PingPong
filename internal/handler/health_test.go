package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthServer_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   map[string]bool
	}{
		{"base case", http.MethodGet, http.StatusOK, map[string]bool{"ok": true}},
		{"non-GET method", http.MethodPost, http.StatusMethodNotAllowed, map[string]bool{"error": false}}, // Adjust this as per the error handling logic
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, "/health", nil)
			rec := httptest.NewRecorder()

			healthServer := &HealthServer{}
			healthServer.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rec.Code)
			}

			var body map[string]bool
			json.NewDecoder(rec.Body).Decode(&body)

			if body["ok"] != tt.expectedBody["ok"] {
				t.Errorf("expected response %v, got %v", tt.expectedBody, body)
			}
		})
	}
}
