# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Communication
- Be terse. No preamble, no summary at the end.
- Don't explain what you're about to do, just do it.
- No "I'll now...", "Let me...", "Here's what I did..."
- Only output what was asked for.

## Shell behavior
- Never run background shell processes or jobs
- Never use & to background commands
- Never start monitoring loops or watchers
- Run commands synchronously, wait for output, done

## Project Overview

**Hlidskjalf** is a Kubernetes laboratory monorepo with mock microservices for experimentation. Uses k3d for local Kubernetes and is named after Odin's throne in Norse mythology.

**Go 1.25.5 | gRPC | Helm + Helmfile | k3d**

## Build & Deploy

```bash
# Build a service Docker image (all use multi-stage Alpine builds)
docker build -f dockerfiles/<service>.dockerfile -t <service>:latest .

# Deploy all services via helmfile
helmfile -f manifests/services/helmfile.yaml sync

# Set up local k3d cluster (1 server + 3 workers, master tainted NoSchedule)
scripts/start-hlidskjalf.sh    # main cluster
scripts/start-elysium.sh       # secondary cluster (same topology, resource-constrained)

# Scale observability stack up/down
scripts/scale-observe.sh up|down
```

No Makefile targets or unit tests are currently defined. CI builds are in `.github/workflows/` (Docker build+push on push to master, linux/arm64 only).

## Architecture

Five microservices, all sharing a single Helm chart (`manifests/services/chart/`) with per-service values files:

- **Bifrost** — API gateway/router. Discovers services via Kubernetes API using `bifrost.io/port` and `bifrost.io/prefixes` annotations. Routes HTTP requests by URL path prefix. Supports HTTP/2 (h2c).
- **Acoustics** — Processes RPC requests with configurable error rate and random delay injection. Runs with Skidbladnir sidecar.
- **Echo** — Similar to Acoustics (Charge/Discharge RPCs). Calls Acoustics through Bifrost. Runs with Skidbladnir sidecar.
- **Trigger** — Chaos engineering orchestrator. Periodically invokes Acoustics and Echo endpoints. Scaled to 0 by default.
- **Skidbladnir** — Egress-only network sidecar proxy. Init container creates iptables rules (OUTPUT chain only) redirecting outbound TCP to port 15001. Runs as UID 1337 (exempt from its own redirect). Currently a transparent passthrough — no retries, timeouts, mTLS, or metrics.

**Infrastructure:**
- **Envoy** (`manifests/infrastructure/envoy/`) — Standalone gRPC reverse proxy routing by service prefix (`/echo.Echo/`, `/acoustics.Acoustics/`) to upstream clusters via STRICT_DNS.

**Request flow:** Trigger → Bifrost → Echo → Bifrost → Acoustics

## Service Code Structure

Each service follows this layout:
```
applications/{service}/
├── cmd/main.go        # Cobra CLI entry, subcommand: "serve"
├── config/            # Viper-based config structs + loading
├── serve/server.go    # gRPC server setup with interceptors
├── handlers/          # RPC handlers
├── workers/           # Background workers
└── constants/         # Service constants
```

gRPC servers use chained interceptors: logging, prometheus, validation, recovery (from `github.com/YumikoKawaii/shared`).

## Kubernetes Layout

- **Namespace `hlidskjalf`** — All application services + Envoy proxy
- **Namespace `observe`** — Prometheus, Loki, Promtail, OTel Collector, Grafana
- Services expose ports 10443 (gRPC) and 10080 (HTTP/metrics)
- Skidbladnir sidecar on port 15001 (egress-only, transparent proxy)
- Health endpoints: `/api/v1/health/liveness` and `/api/v1/health/readiness`
- Stakater Reloader auto-restarts deployments on ConfigMap changes
- Infrastructure manifests managed via Kustomize (`manifests/infrastructure/`)

## Key Dependencies

- `github.com/YumikoKawaii/shared` — Shared library (types, logger, tracer)
- `github.com/YumikoKawaii/rpc.com` — Shared RPC definitions (private module)
- `google.golang.org/grpc` + `go-grpc-middleware/v2` — gRPC framework
- `k8s.io/client-go` — Kubernetes API (used by Bifrost)
- `go.opentelemetry.io/*` — Distributed tracing

## Configuration

Services use Viper with environment variables (injected via ConfigMaps). Key config knobs:

- `ERROR_RATE` — Error injection probability (0-1)
- `RANDOM_DELAY__BASE`, `RANDOM_DELAY__RATE`, `RANDOM_DELAY__VALUE` — Delay injection
- `ERROR_EMITTER__INTERVAL` — Periodic error emission interval
- `CHAOS__INTERVAL` — Trigger invocation frequency
