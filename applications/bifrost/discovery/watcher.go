package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
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
	prefixServiceMap map[string]string   // HTTP/GRPC path prefix → service name
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
		prefixServiceMap: make(map[string]string),
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
				case watch.Modified:
					logger.Infof("[discovery] service modified: %s", service.Name)
					w.deleteService(service.Name)
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

	// guaranteed annotations: prefixes and port
	portValue, f := service.Annotations[constants.PortAnnotation]
	if !f {
		logger.Errorf("[discovery] missing %s for service: %s", constants.PortAnnotation, service.Name)
		return
	}
	port, err := strconv.Atoi(portValue)
	if err != nil {
		logger.Errorf("[discovery] error parsing port value for service %s: %s", service.Name, err.Error())
		return
	}

	prefixesValue, f := service.Annotations[constants.PrefixesAnnotation]
	if !f {
		logger.Errorf("[discovery] missing %s for service: %s", constants.PrefixesAnnotation, service.Name)
		return
	}
	prefixes := make([]string, 0)
	if err := json.Unmarshal([]byte(prefixesValue), &prefixes); err != nil {
		logger.Errorf("[discovery] error parsing prefixes value for service %s: %s", service.Name, err.Error())
		return
	}

	for _, prefix := range prefixes {
		w.prefixServiceMap[prefix] = service.Name
	}

	child, cancel := context.WithCancel(ctx)
	w.servicesMap[service.Name] = NewService(port, cancel)
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
	for prefix, svc := range w.prefixServiceMap {
		if svc == serviceName {
			delete(w.prefixServiceMap, prefix)
			logger.Infof("[discovery] unregistered prefix %s → %s", prefix, serviceName)
		}
	}
	service.Cancel()
	logger.Infof("[discovery] stopped endpoint watcher for %s", serviceName)
}

func (w *Watcher) watchEndpoints(ctx context.Context, service *corev1.Service) {
	for {
		if ctx.Err() != nil {
			logger.Debugf("[discovery] endpoint watcher for %s stopped", service.Name)
			return
		}

		endpoints, err := w.clientset.CoreV1().Endpoints(w.namespace).Get(ctx, service.Name, metav1.GetOptions{})
		if err != nil {
			if ctx.Err() != nil {
				logger.Debugf("[discovery] endpoint watcher for %s stopped", service.Name)
				return
			}
			logger.Infof("[discovery] endpoints for %s not ready, retrying: %v", service.Name, err)
			time.Sleep(2 * time.Second)
			continue
		}
		w.updateServiceEndpoints(service.Name, endpoints)

		endpointsWatcher, err := w.clientset.CoreV1().Endpoints(w.namespace).Watch(ctx, metav1.ListOptions{
			FieldSelector: fmt.Sprintf("metadata.name=%s", service.Name),
		})
		if err != nil {
			if ctx.Err() != nil {
				logger.Debugf("[discovery] endpoint watcher for %s stopped", service.Name)
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
					w.updateServiceEndpoints(service.Name, ep)
				}
			}
		}
	}
}

func (w *Watcher) updateServiceEndpoints(serviceName string, endpoints *corev1.Endpoints) {
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
	svc.UpdateEndpoints(ips)
	logger.Infof("[discovery] %s: %d endpoints %v", serviceName, len(ips), ips)
}

// Resolve finds the backend service for a given request path using longest-prefix match,
// then round-robins across available pod endpoints. Returns (serviceName, target, found)
// where target is "ip:port".
func (w *Watcher) Resolve(path string) (string, string, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	// resolve by prefix
	serviceName := w.resolveByPrefix(path)
	if serviceName == "" {
		return "", "", errors.New("unable to resolve")
	}

	svc, ok := w.servicesMap[serviceName]
	if !ok {
		return "", "", fmt.Errorf("not found service: %s", serviceName)
	}

	return serviceName, svc.GetTarget(), nil
}

// resolveByPrefix finds the service using longest-prefix match against prefixServiceMap.
func (w *Watcher) resolveByPrefix(path string) string {
	var bestPrefix, bestService string
	for prefix, serviceName := range w.prefixServiceMap {
		if strings.HasPrefix(path, prefix) && len(prefix) > len(bestPrefix) {
			bestPrefix = prefix
			bestService = serviceName
		}
	}
	return bestService
}
