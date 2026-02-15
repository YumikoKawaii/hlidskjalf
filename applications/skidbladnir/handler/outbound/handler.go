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

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", h.proxy)
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
