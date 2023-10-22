// health.go
package handler

import (
	"net/http"
	"pingpong/internal/response"
)

// HealthService represents the interface for a health-check service.
type HealthService interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// HealthServer struct implementing the HealthService interface.
type HealthServer struct{}

// ServeHTTP handles the health-check requests.
func (h *HealthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	response.WithJSON(w, http.StatusOK, map[string]bool{"ok": true})
}
