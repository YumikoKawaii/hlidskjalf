# 08 · HPA & metrics-server

**Domain:** Workloads & Scheduling
**Planted faults:** none

## Setup

```bash
./cka/assignments/08-hpa-metrics/setup.sh
```

This creates `autoscale/web` (CPU request 50m, limit 200m) and a Service `web`.

## Task

1. Install metrics-server from its official release manifest and get `kubectl top nodes` working. On a kubeadm cluster it may not work out of the box. Fix it properly and write down why it failed.
2. Create an `autoscaling/v2` HPA for `web`: min 2, max 6, target 60% average CPU utilisation, scale-down stabilisation window 30s.
3. Generate load against `http://web.autoscale` from a busybox pod in the same namespace.
4. Once the HPA has scaled past 2 replicas, save `kubectl -n autoscale get hpa web` to `/opt/cka/hpa.txt`.
5. Stop the load and watch it scale back down to 2.

## Done when

- `kubectl top pods -A` returns data.
- `/opt/cka/hpa.txt` shows more than 2 replicas.
- With the load stopped, `web` settles at 2 replicas.

Assignment 22 reuses metrics-server, so leave it installed.

## Solution

Put your manifests, commands and the metrics-server root cause in `solution/`.
