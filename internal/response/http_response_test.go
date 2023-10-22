package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithError(t *testing.T) {
	rec := httptest.NewRecorder()
	message := "Test error"
	WithError(rec, http.StatusBadRequest, message)

	res := rec.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status code %v; got %v", http.StatusBadRequest, res.StatusCode)
	}

	var body map[string]string
	json.NewDecoder(res.Body).Decode(&body)
	if body["error"] != message {
		t.Errorf("expected error message %q; got %q", message, body["error"])
	}
}

func TestWithJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	data := map[string]string{"key": "value"}
	WithJSON(rec, http.StatusOK, data)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %v; got %v", http.StatusOK, res.StatusCode)
	}

	var body map[string]string
	json.NewDecoder(res.Body).Decode(&body)
	if body["key"] != "value" {
		t.Errorf("expected key value %q; got %q", "value", body["key"])
	}
}
