package handler

import (
	"net/http"

	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/discovery"
	"github.com/YumikoKawaii/shared/logger"
)

type Handler struct {
	watcher *discovery.Watcher
}

func Initialize(watcher *discovery.Watcher) *Handler {
	return &Handler{watcher: watcher}
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

	serviceName, ip, found := h.watcher.Resolve(r.URL.Path)
	if !found {
		logger.Infof("[proxy] no route matched for %s", r.URL.Path)
		http.Error(w, "no route matched", http.StatusNotFound)
		return
	}

	logger.Infof("[proxy] %s %s → %s (%s)", r.Method, r.URL.Path, serviceName, ip)
}
