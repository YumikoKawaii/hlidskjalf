package ratelimit

import (
	"context"
	"path/filepath"
	"time"

	"github.com/YumikoKawaii/shared/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/util/homedir"
)

// ElectionCallbacks defines the callbacks invoked during leader election lifecycle.
type ElectionCallbacks struct {
	OnStartedLeading func(ctx context.Context) // called when this instance becomes leader
	OnStoppedLeading func()                    // called when this instance loses leadership
	OnNewLeader      func(identity string)     // called on all instances when a new leader is elected
}

// buildClientset creates a Kubernetes clientset, trying in-cluster config first
// and falling back to ~/.kube/config for local development. Mirrors the pattern
// used by Bifrost's discovery watcher.
func buildClientset() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	return kubernetes.NewForConfig(config)
}

// RunElection starts a Kubernetes Lease-based leader election. The lease is named
// "<serviceName>-skidbladnir-ratelimit" and scoped to the given namespace.
// Each pod uses its IP as its election identity. This blocks until the context
// is cancelled or the election loop fatally errors.
func RunElection(ctx context.Context, serviceName, namespace, identity string, callbacks ElectionCallbacks) {
	clientset, err := buildClientset()
	if err != nil {
		logger.Fatalf("[ratelimit] failed to build kubernetes clientset: %v", err)
	}

	leaseName := serviceName + "-skidbladnir-ratelimit"

	lock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      leaseName,
			Namespace: namespace,
		},
		Client: clientset.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: identity,
		},
	}

	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:            lock,
		LeaseDuration:   15 * time.Second,
		RenewDeadline:   10 * time.Second,
		RetryPeriod:     2 * time.Second,
		ReleaseOnCancel: true,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: callbacks.OnStartedLeading,
			OnStoppedLeading: callbacks.OnStoppedLeading,
			OnNewLeader:      callbacks.OnNewLeader,
		},
	})
}
