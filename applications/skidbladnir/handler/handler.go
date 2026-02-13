package handler

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/proxy"
	"github.com/YumikoKawaii/shared/logger"
	"golang.org/x/net/http2"
)

type Handler struct {
	h1Transport http.RoundTripper
	h2Transport http.RoundTripper
}

func Initialize() *Handler {
	return &Handler{
		h1Transport: http.DefaultTransport,
		h2Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
				var d net.Dialer
				return d.DialContext(ctx, network, addr)
			},
		},
	}
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
	dst, ok := proxy.OriginalDstFromContext(r.Context())
	if !ok {
		http.Error(w, "original destination unavailable", http.StatusBadGateway)
		return
	}

	var transport http.RoundTripper
	if r.ProtoMajor == 2 {
		transport = h.h2Transport
	} else {
		transport = h.h1Transport
	}

	logger.Infof("[proxy] %s %s → %s proto=%s", r.Method, r.URL.Path, dst, r.Proto)

	rp := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = dst
			req.Host = ""
		},
		Transport: transport,
	}
	rp.ServeHTTP(w, r)
}
