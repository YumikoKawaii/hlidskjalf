package handler

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"
	"time"

	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/constants"
	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/discovery"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/http2"
)

var proxyRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "bifrost_proxy_requests_total",
		Help: "Total number of proxied requests",
	},
	[]string{"service", "method"},
)

type Handler struct {
	watcher     *discovery.Watcher
	h1Transport http.RoundTripper
	h2Transport http.RoundTripper
	proxies     sync.Map // "h1:ip:port" or "h2:ip:port" → *httputil.ReverseProxy
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
		if strings.HasPrefix(r.URL.Path, "/api/v1/health/") || r.URL.Path == "/metrics" || strings.HasPrefix(r.URL.Path, "/debug/pprof/") {
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

	logger.Debugf("[proxy] %s %s → %s (%s)", r.Method, r.URL.Path, service, target)

	proxyRequestsTotal.WithLabelValues(service, r.Method).Inc()
	h.resolveProxy(target, r.ProtoMajor).ServeHTTP(w, r)
}

// resolveProxy returns a cached ReverseProxy for the given target and protocol,
// creating one on first access. Uses sync.Map for lock-free reads on the hot path.
func (h *Handler) resolveProxy(target string, protoMajor int) *httputil.ReverseProxy {
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
			req.URL.Scheme = constants.HTTPScheme
			req.URL.Host = target
			req.Host = ""
		},
		Transport: transport,
	}

	actual, _ := h.proxies.LoadOrStore(key, proxy)
	return actual.(*httputil.ReverseProxy)
}
