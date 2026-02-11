package discovery

import (
	"context"
	"fmt"
	"sync"
	"time"

	"path/filepath"

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
	serviceIpsMap    map[string][]string // service name → list of pod IPs
	serviceCancelMap map[string]context.CancelFunc
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
		serviceIpsMap:    make(map[string][]string),
		serviceCancelMap: make(map[string]context.CancelFunc),
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
	if _, exists := w.serviceCancelMap[service.Name]; exists {
		logger.Infof("[discovery] service %s already watched, skipping", service.Name)
		return
	}
	child, cancel := context.WithCancel(ctx)
	w.serviceCancelMap[service.Name] = cancel
	logger.Infof("[discovery] starting endpoint watcher for %s", service.Name)
	go w.watchEndpoints(child, service)
}

func (w *Watcher) deleteService(serviceName string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	delete(w.serviceIpsMap, serviceName)
	if cancel, ok := w.serviceCancelMap[serviceName]; ok {
		delete(w.serviceCancelMap, serviceName)
		cancel()
		logger.Infof("[discovery] stopped endpoint watcher for %s", serviceName)
	} else {
		logger.Infof("[discovery] no watcher found for %s, skipping", serviceName)
	}
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
	w.mu.Lock()
	defer w.mu.Unlock()
	w.serviceIpsMap[serviceName] = ips
	logger.Infof("[discovery] %s: %d endpoints %v", serviceName, len(ips), ips)
}
