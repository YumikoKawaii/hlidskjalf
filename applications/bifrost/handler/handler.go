package handler

import (
	"net/http"

	"github.com/YumikoKawaii/shared/logger"
)

type Handler struct {
}

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
	logger.Infof("received: %s %s proto=%s content-type=%s", r.Method, r.URL.Path, r.Proto, r.Header.Get("Content-Type"))
}
