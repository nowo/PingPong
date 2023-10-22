package handler

import (
	"net/http"
	"pingpong/internal/response"
)

type PingService interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type PingServer struct{}

func (p *PingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	response.WithJSON(w, http.StatusOK, map[string]string{"response": "pong"})
}
