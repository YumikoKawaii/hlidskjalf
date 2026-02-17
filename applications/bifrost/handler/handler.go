package handler

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/constants"
	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/discovery"
	"github.com/YumikoKawaii/shared/logger"
	"golang.org/x/net/http2"
)

type Handler struct {
	watcher     *discovery.Watcher
	h1Transport http.RoundTripper
	h2Transport http.RoundTripper
}

func Initialize(watcher *discovery.Watcher) *Handler {
	return &Handler{
		watcher: watcher,
		h1Transport: &http.Transport{
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     90 * time.Second,
		},
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
	logger.Infof("received: %s %s proto=%s content-type=%s", r.Method, r.URL.Path, r.Proto, r.Header.Get("Content-Type"))

	service, target, err := h.watcher.Resolve(r.URL.Path)
	if err != nil {
		logger.Infof("[proxy] resolve failed for %s: %v", r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var transport http.RoundTripper
	if r.ProtoMajor == 2 {
		transport = h.h2Transport
	} else {
		transport = h.h1Transport
	}

	logger.Infof("[proxy] %s %s → %s (%s)", r.Method, r.URL.Path, service, target)

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = constants.HTTPScheme
			req.URL.Host = target
			req.Host = ""
		},
		Transport: transport,
	}
	proxy.ServeHTTP(w, r)
}
