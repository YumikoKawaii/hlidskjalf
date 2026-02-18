package discovery

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"sync"
	"time"

	"github.com/YumikoKawaii/shared/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Resolver struct {
	clientset *kubernetes.Clientset
	namespace string
	ctx       context.Context

	mu       sync.RWMutex
	services map[string]*Service
}

func NewResolver(namespace string) (*Resolver, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("failed to get kubernetes config: %w", err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	return &Resolver{
		clientset: clientset,
		namespace: namespace,
		services:  make(map[string]*Service),
	}, nil
}

// Start stores the context for lazy endpoint watchers.
func (r *Resolver) Start(ctx context.Context) {
	r.ctx = ctx
}

// Resolve looks up the hostname from host (e.g. "acoustics:10443") and
// returns a round-robin pod IP with the original port. On first access for
// a hostname, it lazily starts an endpoint watcher. Returns ok=false when
// no endpoints are available yet (falls back to passthrough).
func (r *Resolver) Resolve(host string) (string, bool) {
	hostname, port, err := net.SplitHostPort(host)
	if err != nil {
		hostname = host
		port = ""
	}

	svc := r.getOrWatch(hostname)

	if port == "" {
		port = "80"
	}

	target, err := svc.GetTarget(port)
	if err != nil {
		return "", false
	}
	return target, true
}

// getOrWatch returns the Service for hostname, lazily starting a watcher if needed.
func (r *Resolver) getOrWatch(hostname string) *Service {
	r.mu.RLock()
	svc, ok := r.services[hostname]
	r.mu.RUnlock()
	if ok {
		return svc
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	// double-check after acquiring write lock
	if svc, ok = r.services[hostname]; ok {
		return svc
	}

	svc = NewService()
	r.services[hostname] = svc
	go r.watchEndpoints(r.ctx, hostname)
	logger.Infof("[balancer] started watching endpoints for %s", hostname)
	return svc
}

func (r *Resolver) watchEndpoints(ctx context.Context, serviceName string) {
	for {
		if ctx.Err() != nil {
			logger.Debugf("[balancer] endpoint watcher for %s stopped", serviceName)
			return
		}

		endpoints, err := r.clientset.CoreV1().Endpoints(r.namespace).Get(ctx, serviceName, metav1.GetOptions{})
		if err != nil {
			if ctx.Err() != nil {
				logger.Debugf("[balancer] endpoint watcher for %s stopped", serviceName)
				return
			}
			logger.Infof("[balancer] endpoints for %s not ready, retrying: %v", serviceName, err)
			time.Sleep(2 * time.Second)
			continue
		}
		r.updateEndpoints(serviceName, endpoints)

		endpointsWatcher, err := r.clientset.CoreV1().Endpoints(r.namespace).Watch(ctx, metav1.ListOptions{
			FieldSelector: fmt.Sprintf("metadata.name=%s", serviceName),
		})
		if err != nil {
			if ctx.Err() != nil {
				logger.Debugf("[balancer] endpoint watcher for %s stopped", serviceName)
				return
			}
			logger.Infof("[balancer] watch for %s failed, retrying: %v", serviceName, err)
			time.Sleep(2 * time.Second)
			continue
		}

		for event := range endpointsWatcher.ResultChan() {
			switch event.Type {
			case watch.Added, watch.Modified:
				if ep, ok := event.Object.(*corev1.Endpoints); ok {
					r.updateEndpoints(serviceName, ep)
				}
			}
		}
	}
}

func (r *Resolver) updateEndpoints(serviceName string, endpoints *corev1.Endpoints) {
	ips := make([]string, 0)
	for _, subset := range endpoints.Subsets {
		for _, addr := range subset.Addresses {
			ips = append(ips, addr.IP)
		}
	}

	r.mu.RLock()
	svc, ok := r.services[serviceName]
	r.mu.RUnlock()
	if !ok {
		return
	}
	svc.UpdateEndpoints(ips)
	logger.Infof("[balancer] %s: %d endpoints %v", serviceName, len(ips), ips)
}
