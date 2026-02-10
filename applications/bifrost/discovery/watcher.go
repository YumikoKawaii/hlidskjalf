package discovery

import (
	"context"
	"fmt"
	"sync"

	"github.com/YumikoKawaii/shared/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type Watcher struct {
	clientset *kubernetes.Clientset
	namespace string

	mu       sync.RWMutex
	services map[string][]string // service name → list of pod IPs
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
		clientset: clientset,
		namespace: namespace,
		services:  make(map[string][]string),
	}, nil
}

// Watch starts watching endpoints for a given service name.
// It blocks, so call it in a goroutine.
func (w *Watcher) Watch(ctx context.Context, serviceName string) {
	for {
		// Initial list to populate
		endpoints, err := w.clientset.CoreV1().Endpoints(w.namespace).Get(ctx, serviceName, metav1.GetOptions{})
		if err != nil {
			logger.Errorf("failed to get endpoints for %s: %v", serviceName, err)
			return
		}
		w.updateAddresses(serviceName, endpoints)

		// Start watching for changes
		watcher, err := w.clientset.CoreV1().Endpoints(w.namespace).Watch(ctx, metav1.ListOptions{
			FieldSelector: fmt.Sprintf("metadata.name=%s", serviceName),
		})
		if err != nil {
			logger.Errorf("failed to watch endpoints for %s: %v", serviceName, err)
			return
		}

		for event := range watcher.ResultChan() {
			switch event.Type {
			case watch.Added, watch.Modified:
				if ep, ok := event.Object.(*corev1.Endpoints); ok {
					w.updateAddresses(serviceName, ep)
				}
			case watch.Deleted:
				w.mu.Lock()
				delete(w.services, serviceName)
				w.mu.Unlock()
				logger.Infof("[ディスカバリ] %s: endpoints deleted", serviceName)
			}
		}

		// Watcher closed (e.g. timeout), restart
		logger.Infof("[ディスカバリ] %s: watcher closed, restarting...", serviceName)
	}
}

// GetAddresses returns the current list of pod IPs for a service.
func (w *Watcher) GetAddresses(serviceName string) []string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.services[serviceName]
}

func (w *Watcher) updateAddresses(serviceName string, endpoints *corev1.Endpoints) {
	var addrs []string
	for _, subset := range endpoints.Subsets {
		for _, addr := range subset.Addresses {
			addrs = append(addrs, addr.IP)
		}
	}

	w.mu.Lock()
	w.services[serviceName] = addrs
	w.mu.Unlock()

	logger.Infof("[ディスカバリ] %s: %d endpoints %v", serviceName, len(addrs), addrs)
}
