// Package outbound implements the outbound egress proxy for Skidbladnir.
// This handles traffic intercepted by iptables OUTPUT chain rules.
package outbound

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/discovery"
	"github.com/YumikoKawaii/shared/logger"
	"golang.org/x/net/http2"
)

type Handler struct {
	resolver    *discovery.Resolver
	h1Transport http.RoundTripper
	h2Transport http.RoundTripper
	proxies     sync.Map // "h1:ip:port" or "h2:ip:port" → *httputil.ReverseProxy
}

func Initialize(resolver *discovery.Resolver) *Handler {
	return &Handler{
		resolver: resolver,
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

	// Resolve to pod IP if balancer is configured for this service
	target := host
	if h.resolver != nil {
		if resolved, ok := h.resolver.Resolve(host); ok {
			target = resolved
		}
	}

	logger.Debugf("[skidbladnir] %s %s → %s (host=%s) proto=%s", r.Method, r.RequestURI, target, host, r.Proto)

	h.resolveProxy(target, host, r.ProtoMajor).ServeHTTP(w, r)
}

// resolveProxy returns a cached ReverseProxy for the given target and protocol,
// creating one on first access. Uses sync.Map for lock-free reads on the hot path.
func (h *Handler) resolveProxy(target, originalHost string, protoMajor int) *httputil.ReverseProxy {
	var key string
	var transport http.RoundTripper
	if protoMajor == 2 {
		key = fmt.Sprintf("h2:%s", target)
		transport = h.h2Transport
	} else {
		key = fmt.Sprintf("h1:%s", target)
		transport = h.h1Transport
	}

	if v, ok := h.proxies.Load(key); ok {
		return v.(*httputil.ReverseProxy)
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.Host = originalHost
		},
		Transport: transport,
	}

	actual, _ := h.proxies.LoadOrStore(key, proxy)
	return actual.(*httputil.ReverseProxy)
}
