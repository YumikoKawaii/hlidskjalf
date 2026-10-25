# 14 · Gateway API

**Domain:** Services & Networking
**Planted faults:** none

## Setup

```bash
./cka/assignments/14-gateway-api/setup.sh
```

This creates `gw/app-v1` and `gw/app-v2` with matching Services. They answer `v1` and `v2`.

## Task

1. Install the Gateway API **standard channel** CRDs from the latest release.
2. Install a Gateway controller of your choice (Envoy Gateway, NGINX Gateway Fabric, Cilium, …). The lab has no LoadBalancer, so expose the gateway's data plane on a NodePort.
3. In `gw`, create a Gateway `web-gw` with an HTTP listener on port 80 that accepts routes from its own namespace.
4. In `gw`, create an HTTPRoute `web-route` for host `web.lab.local`:
   - `/v1` → `app-v1`
   - `/v2` → `app-v2`
   - `/canary` → 80% `app-v1`, 20% `app-v2`
5. From `cp1`, send 20 requests to `/canary` with `Host: web.lab.local` and save the responses to `/opt/cka/canary.txt`.

## Done when

- `kubectl -n gw get gateway web-gw` shows `PROGRAMMED=True`.
- `kubectl -n gw get httproute web-route -o yaml` shows `Accepted` and `ResolvedRefs` both `True`.
- `/v1` returns `v1`, `/v2` returns `v2`, and `/opt/cka/canary.txt` is mostly `v1`.

## Solution

Put your controller install steps and manifests in `solution/`.
