package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"path/filepath"

	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/constants"
	"github.com/YumikoKawaii/shared/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Watcher struct {
	clientset *kubernetes.Clientset
	namespace string

	mu               sync.RWMutex
	servicesMap      map[string]*Service // service name → Service
	serviceRoutesMap map[string]string   // HTTP path prefix → service name
}

func NewWatcher(namespace string) (*Watcher, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fall back to kubeconfig for local development
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

	return &Watcher{
		clientset:        clientset,
		namespace:        namespace,
		servicesMap:      make(map[string]*Service),
		serviceRoutesMap: make(map[string]string),
	}, nil
}

func (w *Watcher) DiscoverServices(ctx context.Context) {
	for {
		services, err := w.clientset.CoreV1().Services(w.namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			logger.Errorf("error discover services: %s", err.Error())
			return
		}
		logger.Infof("[discovery] found %d services", len(services.Items))
		// for each service, initialize watcher for each services
		for _, service := range services.Items {
			w.addService(ctx, &service)
		}

		// watching for services
		serviceWatcher, err := w.clientset.CoreV1().Services(w.namespace).Watch(ctx, metav1.ListOptions{
			ResourceVersion: services.ResourceVersion,
		})
		if err != nil {
			logger.Errorf("error watching services: %s", err.Error())
			return
		}

		for event := range serviceWatcher.ResultChan() {
			service, ok := event.Object.(*corev1.Service)
			if ok {
				switch event.Type {
				case watch.Added:
					logger.Infof("[discovery] service added: %s", service.Name)
					w.addService(ctx, service)
				case watch.Deleted:
					logger.Infof("[discovery] service deleted: %s", service.Name)
					w.deleteService(service.Name)
				}
			}
		}
	}
}

func (w *Watcher) addService(ctx context.Context, service *corev1.Service) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if _, exists := w.servicesMap[service.Name]; exists {
		logger.Infof("[discovery] service %s already watched, skipping", service.Name)
		return
	}

	// Skip services not managed by bifrost
	if service.Annotations[constants.Enable] != "true" {
		logger.Infof("[discovery] service %s has no %s annotation, skipping", service.Name, constants.Enable)
		return
	}

	// Optionally register HTTP/1.1 routes
	if routeAnnotation, ok := service.Annotations[constants.HTTPRoutes]; ok {
		var routes []string
		if err := json.Unmarshal([]byte(routeAnnotation), &routes); err != nil {
			logger.Errorf("[discovery] failed to parse %s annotation for %s: %v", constants.HTTPRoutes, service.Name, err)
		} else {
			for _, route := range routes {
				w.serviceRoutesMap[route] = service.Name
				logger.Infof("[discovery] registered HTTP route %s → %s", route, service.Name)
			}
		}
	}

	child, cancel := context.WithCancel(ctx)
	w.servicesMap[service.Name] = NewService(service.Name, cancel)
	logger.Infof("[discovery] starting endpoint watcher for %s", service.Name)
	go w.watchEndpoints(child, service)
}

func (w *Watcher) deleteService(serviceName string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	service, f := w.servicesMap[serviceName]
	if !f {
		logger.Infof("[discovery] no metadata found for %s, skipping", serviceName)
		return
	}
	delete(w.servicesMap, serviceName)
	// Remove route entries for this service
	for route, svc := range w.serviceRoutesMap {
		if svc == serviceName {
			delete(w.serviceRoutesMap, route)
			logger.Infof("[discovery] unregistered HTTP route %s → %s", route, serviceName)
		}
	}
	service.Cancel()
	logger.Infof("[discovery] stopped endpoint watcher for %s", serviceName)
}

func (w *Watcher) watchEndpoints(ctx context.Context, service *corev1.Service) {
	for {
		if ctx.Err() != nil {
			logger.Infof("[discovery] endpoint watcher for %s stopped", service.Name)
			return
		}

		endpoints, err := w.clientset.CoreV1().Endpoints(w.namespace).Get(ctx, service.Name, metav1.GetOptions{})
		if err != nil {
			if ctx.Err() != nil {
				logger.Infof("[discovery] endpoint watcher for %s stopped", service.Name)
				return
			}
			logger.Infof("[discovery] endpoints for %s not ready, retrying: %v", service.Name, err)
			time.Sleep(2 * time.Second)
			continue
		}
		w.updateServiceIps(service.Name, endpoints)

		endpointsWatcher, err := w.clientset.CoreV1().Endpoints(w.namespace).Watch(ctx, metav1.ListOptions{
			FieldSelector: fmt.Sprintf("metadata.name=%s", service.Name),
		})
		if err != nil {
			if ctx.Err() != nil {
				logger.Infof("[discovery] endpoint watcher for %s stopped", service.Name)
				return
			}
			logger.Infof("[discovery] watch for %s failed, retrying: %v", service.Name, err)
			time.Sleep(2 * time.Second)
			continue
		}

		for event := range endpointsWatcher.ResultChan() {
			switch event.Type {
			case watch.Added, watch.Modified:
				if ep, ok := event.Object.(*corev1.Endpoints); ok {
					w.updateServiceIps(service.Name, ep)
				}
			}
		}
	}
}

func (w *Watcher) updateServiceIps(serviceName string, endpoints *corev1.Endpoints) {
	ips := make([]string, 0)
	for _, subset := range endpoints.Subsets {
		for _, addr := range subset.Addresses {
			ips = append(ips, addr.IP)
		}
	}
	w.mu.RLock()
	svc, ok := w.servicesMap[serviceName]
	w.mu.RUnlock()
	if !ok {
		return
	}
	svc.UpdateIPs(ips)
	logger.Infof("[discovery] %s: %d endpoints %v", serviceName, len(ips), ips)
}

// Resolve finds the backend service for a given request path,
// then round-robins across available pod IPs. Returns (serviceName, ip, found).
//
// For HTTP/2.0 (gRPC): parses the service name from the path (/{package}.{Service}/{Method})
// and does a direct lookup in servicesMap.
// For HTTP/1.1: uses longest-prefix match against serviceRoutesMap.
func (w *Watcher) Resolve(path string, proto string) (string, string, bool) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	var serviceName string
	if proto == "HTTP/2.0" {
		serviceName = resolveGRPCService(path)
	} else {
		serviceName = w.resolveHTTPService(path)
	}
	if serviceName == "" {
		return "", "", false
	}

	svc, ok := w.servicesMap[serviceName]
	if !ok {
		return serviceName, "", false
	}

	ip, ok := svc.GetTarget()
	if !ok {
		return serviceName, "", false
	}
	return serviceName, ip, true
}

// resolveGRPCService extracts the k8s service name from a gRPC path.
// Path format: /{package}.{Service}/{Method} → returns {package}.
func resolveGRPCService(path string) string {
	// Trim leading "/"
	trimmed := strings.TrimPrefix(path, "/")
	// Get the "{package}.{Service}" segment before the method
	slashIdx := strings.Index(trimmed, "/")
	if slashIdx < 0 {
		return ""
	}
	fullService := trimmed[:slashIdx]
	// Extract package name (before the first dot)
	dotIdx := strings.Index(fullService, ".")
	if dotIdx < 0 {
		return ""
	}
	return fullService[:dotIdx]
}

// resolveHTTPService finds the service for an HTTP/1.1 path using longest-prefix match.
func (w *Watcher) resolveHTTPService(path string) string {
	var bestPrefix, bestService string
	for prefix, serviceName := range w.serviceRoutesMap {
		if strings.HasPrefix(path, prefix) && len(prefix) > len(bestPrefix) {
			bestPrefix = prefix
			bestService = serviceName
		}
	}
	return bestService
}
