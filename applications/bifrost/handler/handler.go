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

func (h *Handler) Entry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("received: %s %s proto=%s content-type=%s", r.Method, r.URL.Path, r.Proto, r.Header.Get("Content-Type"))
	}
}
