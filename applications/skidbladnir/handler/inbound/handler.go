// Package inbound implements a reverse proxy for inbound traffic interception.
// Skidbladnir intercepts inbound traffic via iptables PREROUTING rules and
// applies distributed rate limiting before forwarding to the local application.
package inbound

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/constants"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/limiter"
	"github.com/YumikoKawaii/shared/logger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Handler proxies inbound traffic to a local application port after rate limiting.
type Handler struct {
	targetPort  int
	h1Transport http.RoundTripper
	h2Transport http.RoundTripper
	rlManager   *limiter.Manager
}

// NewHandler creates an inbound proxy handler that forwards to the given local port.
func NewHandler(targetPort int, rlManager *limiter.Manager) *Handler {
	return &Handler{
		targetPort:  targetPort,
		h1Transport: http.DefaultTransport,
		h2Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
				var d net.Dialer
				return d.DialContext(ctx, network, addr)
			},
		},
		rlManager: rlManager,
	}
}

// ListenAndServe starts the inbound proxy listener with h2c support.
func (h *Handler) ListenAndServe(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.proxy)

	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(mux, h2s),
	}

	logger.Infof("[skidbladnir] inbound proxy listening on %s → localhost:%d", addr, h.targetPort)
	return server.ListenAndServe()
}

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request) {
	// Apply rate limiting on inbound traffic, skip health probes
	if !strings.HasPrefix(r.URL.Path, "/api/v1/health/") && h.rlManager != nil && !h.rlManager.Allow() {
		logger.Warnf("[skidbladnir] inbound rate limit exceeded: %s %s", r.Method, r.RequestURI)
		http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	target := fmt.Sprintf("127.0.0.1:%d", h.targetPort)

	var transport http.RoundTripper
	if r.ProtoMajor == 2 {
		transport = h.h2Transport
	} else {
		transport = h.h1Transport
	}

	logger.Debugf("[skidbladnir] inbound %s %s → %s proto=%s", r.Method, r.RequestURI, target, r.Proto)

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = constants.HTTPScheme
			req.URL.Host = target
			req.Host = r.Host
		},
		Transport: transport,
	}
	proxy.ServeHTTP(w, r)
}
