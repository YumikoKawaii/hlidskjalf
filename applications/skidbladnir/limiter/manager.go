package limiter

import (
	"context"
	"fmt"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/constants"
	"github.com/YumikoKawaii/shared/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/util/homedir"
)

// Manager is the top-level coordinator for distributed rate limiting.
// Created by the server when rate limiting is enabled, then passed to handlers.
//
// Every pod runs an Operator for local rate-limiting decisions.
// The pod that wins the Lease election also runs a Coordinator that computes
// fair-share allocations for all operators (including itself).
// Allow() always goes through the Operator — no special-casing for the leader.
type Manager struct {
	cfg *config.Application

	coordinator atomic.Pointer[Coordinator] // non-nil only while this pod holds the lease
	operator    *Operator                   // always active on every pod

	k8sClient *kubernetes.Clientset // used for Lease-based leader election
}

func NewManager(cfg *config.Application) (*Manager, error) {
	clusterCfg, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		clusterCfg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	k8sClient, err := kubernetes.NewForConfig(clusterCfg)
	if err != nil {
		return nil, err
	}

	ttl := time.Duration(cfg.Limiter.TTL) * time.Millisecond
	interval := time.Duration(cfg.Limiter.Interval) * time.Millisecond

	return &Manager{
		cfg:       cfg,
		operator:  NewOperator(cfg.Ip, cfg.Limiter.LeaderPort, interval, ttl),
		k8sClient: k8sClient,
	}, nil
}

// Allow checks whether the current request should be permitted.
// Always delegates to the Operator's local token bucket — same code path
// on every pod regardless of leader/follower role.
func (m *Manager) Allow() bool {
	allowed := m.operator.Allow()
	logger.Infof("[limiter] allow=%v tokens=%.1f", allowed, m.operator.limiter.Tokens())
	return allowed
}

// Start launches the leader election goroutine and the operator fetch loop.
// On election win, a Coordinator is created with the global budget.
// On election loss, the Coordinator is cleared.
// The operator fetch loop runs on every pod, requesting capacity from
// whichever pod currently holds the lease.
func (m *Manager) Start(ctx context.Context) {
	go m.Election(ctx)
	go m.operator.Fetch(ctx)
}

// Election runs the Kubernetes Lease-based leader election loop.
// Blocks until the context is cancelled. Only one pod per service holds the
// lease at a time; the winner creates a Coordinator with the full global budget.
func (m *Manager) Election(ctx context.Context) {
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock: &resourcelock.LeaseLock{
			LeaseMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", m.cfg.Service, constants.Identity, constants.Limiter),
				Namespace: m.cfg.Namespace,
			},
			Client: m.k8sClient.CoordinationV1(),
			LockConfig: resourcelock.ResourceLockConfig{
				Identity: m.cfg.Ip,
			},
		},
		LeaseDuration:   15 * time.Second,
		RenewDeadline:   10 * time.Second,
		RetryPeriod:     2 * time.Second,
		ReleaseOnCancel: true,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: m.onStartedLeading,
			OnStoppedLeading: m.onStoppedLeading,
			OnNewLeader:      m.onNewLeader,
		},
	})
}

func (m *Manager) onStartedLeading(ctx context.Context) {
	logger.Infof("[limiter] became coordinator (identity=%s)", m.cfg.Ip)
	ttl := time.Duration(m.cfg.Limiter.TTL) * time.Millisecond
	m.coordinator.Store(NewCoordinator(m.cfg.Limiter.RPS, m.cfg.Limiter.Burst, ttl))
	<-ctx.Done() // block until leadership is lost
}

func (m *Manager) onStoppedLeading() {
	logger.Info("[limiter] stopped coordinating")
	m.coordinator.Store(nil)
}

func (m *Manager) onNewLeader(identity string) {
	logger.Infof("[limiter] new coordinator: %s", identity)
	if err := m.operator.SetLeader(identity); err != nil {
		logger.Errorf("[limiter] error connecting to coordinator: %s", err.Error())
	}
}

// Capacity returns the fair-share allocation for an operator. Called by the gRPC handler.
// Only the coordinator (leader pod) can grant capacity; non-leaders return an error.
func (m *Manager) Capacity(operatorIP string) (Capacity, error) {
	if c := m.coordinator.Load(); c != nil {
		return c.Capacity(operatorIP), nil
	}
	return Capacity{}, fmt.Errorf("not the coordinator")
}

// Unregister removes an operator from the coordinator's tracking map immediately.
// Called by the gRPC handler on graceful shutdown.
func (m *Manager) Unregister(operatorIP string) error {
	if c := m.coordinator.Load(); c != nil {
		c.Unregister(operatorIP)
		return nil
	}
	return fmt.Errorf("not the coordinator")
}

// Stop performs graceful shutdown: notifies the coordinator that this operator
// is leaving so it can immediately rebalance fair-share across remaining operators.
func (m *Manager) Stop() {
	if m.operator != nil {
		m.operator.Unregister()
	}
}

// GetStatus returns the current state for the debug/status endpoint.
func (m *Manager) GetStatus() Status {
	leaderIP := ""
	if v := m.operator.leaderIP.Load(); v != nil {
		leaderIP, _ = v.(string)
	}
	return Status{
		LeaderIP:    leaderIP,
		PodIP:       m.cfg.Ip,
		GlobalRps:   m.cfg.Limiter.RPS,
		GlobalBurst: int32(m.cfg.Limiter.Burst),
	}
}
