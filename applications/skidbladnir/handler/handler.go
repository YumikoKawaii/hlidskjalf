package handler

import (
	"net/http"
)

type Handler struct{}

func Initialize() *Handler {
	return &Handler{}
}

func (h *Handler) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health/liveness", h.healthCheck)
	mux.HandleFunc("/api/v1/health/readiness", h.healthCheck)
	mux.HandleFunc("/", h.proxy)
	return mux
}

func (h *Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not yet implemented", http.StatusBadGateway)
}
