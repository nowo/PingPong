package response

import (
	"encoding/json"
	"net/http"
)

// WithError sends an HTTP response with a specific error message and status code.
func WithError(w http.ResponseWriter, code int, message string) {
	WithJSON(w, code, map[string]string{"error": message})
}

// WithJSON sends an HTTP response with a given status code and payload.
func WithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
