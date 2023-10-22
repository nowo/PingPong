package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingServer_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   map[string]string
	}{
		{"base case", http.MethodGet, http.StatusOK, map[string]string{"response": "pong"}},
		{"non-GET method", http.MethodPost, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, "/ping", nil)
			rec := httptest.NewRecorder()

			pingServer := &PingServer{}
			pingServer.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rec.Code)
			}

			var body map[string]string
			json.NewDecoder(rec.Body).Decode(&body)

			if body["response"] != tt.expectedBody["responseeeeee"] {
				t.Errorf("expected response %v, got %v", tt.expectedBody, body)
			}
		})
	}
}
