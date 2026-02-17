<div align="center">

# Skidbladnir

### *The Folding Ship — Always Finds Favorable Wind*

---

**Egress Sidecar Proxy**

</div>

---

## About

> *"Skidbladnir is the best of ships, crafted with cunning..."*

**Skidbladnir** is a transparent egress proxy that rides alongside application containers as a sidecar. An init container sets up iptables rules to intercept all outbound TCP traffic and redirect it through the proxy.

It speaks both HTTP/1.1 and h2c, matching the protocol of whatever passes through.

---

## How It Works

The init container creates iptables rules on the OUTPUT chain:

- Traffic from UID 1337 (Skidbladnir itself) passes through untouched
- Loopback traffic is excluded
- Everything else is redirected to port 15001

Skidbladnir reads the `Host` header and forwards the request to its original destination.

### Distributed Rate Limiting

When enabled, Skidbladnir instances of the same service coordinate to enforce a global rate limit using a Coordinator/Operator model:

- **Coordinator** — Runs only on the pod that holds the Kubernetes Lease. Tracks all operators and computes fair-share allocations (global RPS / operator count).
- **Operator** — Runs on every pod (including the leader). Owns a local token bucket checked synchronously on every request. Periodically fetches its fair-share capacity from the Coordinator via gRPC.

Leader election uses `coordination.k8s.io/v1` Leases. When a pod shuts down, it unregisters from the Coordinator so rebalancing happens immediately without waiting for TTL expiry.

Requests that exceed the rate limit receive HTTP 429. Health probe paths (`/api/v1/health/*`) are excluded from rate limiting.

---

## Configuration

| Variable | Default | Description |
|---|---|---|
| `SKIDBLADNIR_SERVER__HTTP` | `0.0.0.0:15001` | Outbound proxy listen address |
| `SKIDBLADNIR_SERVER__GRPC` | `0.0.0.0:15443` | gRPC server (Coordinator RPCs) |
| `SERVICE` | — | Service name (used for Lease identity) |
| `NAMESPACE` | — | Kubernetes namespace |
| `IP` | — | Pod IP (injected via downward API) |
| `LIMITER__ENABLED` | `false` | Enable distributed rate limiting |
| `LIMITER__RPS` | `100` | Global RPS cap across all pods |
| `LIMITER__BURST` | `100` | Global burst cap across all pods |
| `LIMITER__LEADER_PORT` | `15443` | gRPC port to reach the Coordinator |
| `LIMITER__TTL` | `2000` | Operator eviction TTL (ms) |
| `LIMITER__INTERVAL` | `500` | Operator capacity fetch interval (ms) |

---

<div align="center">

*~ Small enough to fold, swift enough to carry ~*

</div>
