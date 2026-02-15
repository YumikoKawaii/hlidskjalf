// Package outbound implements the outbound egress proxy for Skidbladnir.
// This handles traffic intercepted by iptables OUTPUT chain rules.
package outbound

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"

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

// Register wraps the given mux so that internal endpoints (health, metrics,
// rate limiter API) are served by the mux, while all other requests are
// forwarded by the egress proxy.
//
// The shared server registers a catch-all "/" on the mux (grpc-gateway),
// so we can't distinguish by pattern match alone. Instead we check if the
// request's Host targets the listener itself (internal) vs an external
// destination (iptables-redirected egress traffic).
func (h *Handler) Register(mux *http.ServeMux, listenAddr string) http.Handler {
	_, listenPort, _ := net.SplitHostPort(listenAddr)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Requests targeting our own listener port are internal (health probes,
		// metrics scrapes, rate limiter API). Egress-redirected traffic has the
		// original destination host which won't match our listener.
		_, reqPort, _ := net.SplitHostPort(r.Host)
		if reqPort == listenPort {
			mux.ServeHTTP(w, r)
			return
		}
		h.proxy(w, r)
	})
}

func (h *Handler) proxy(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	if host == "" {
		http.Error(w, "missing Host header", http.StatusBadRequest)
		return
	}

	var transport http.RoundTripper
	if r.ProtoMajor == 2 {
		transport = h.h2Transport
	} else {
		transport = h.h1Transport
	}

	logger.Infof("[skidbladnir] %s %s → %s proto=%s", r.Method, r.RequestURI, host, r.Proto)

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = host
			req.Host = host
		},
		Transport: transport,
	}
	proxy.ServeHTTP(w, r)
}
