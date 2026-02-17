package handler

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/config"
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

func Initialize(watcher *discovery.Watcher, transport *config.TransportConfig) *Handler {
	return &Handler{
		watcher: watcher,
		h1Transport: &http.Transport{
			MaxIdleConns:        transport.MaxIdleConns,
			MaxIdleConnsPerHost: transport.MaxIdleConnsPerHost,
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

// Wrap returns an http.Handler that routes health and metrics to the shared mux,
// and everything else to the reverse proxy.
func (h *Handler) Wrap(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/v1/health/") || r.URL.Path == "/metrics" {
			mux.ServeHTTP(w, r)
			return
		}
		h.proxy(w, r)
	})
}

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request) {
	logger.Debugf("received: %s %s proto=%s content-type=%s", r.Method, r.URL.Path, r.Proto, r.Header.Get("Content-Type"))

	service, target, err := h.watcher.Resolve(r.URL.Path)
	if err != nil {
		logger.Debugf("[proxy] resolve failed for %s: %v", r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var transport http.RoundTripper
	if r.ProtoMajor == 2 {
		transport = h.h2Transport
	} else {
		transport = h.h1Transport
	}

	logger.Debugf("[proxy] %s %s → %s (%s)", r.Method, r.URL.Path, service, target)

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
