# 20 · Broken Service

**Domain:** Troubleshooting
**Planted faults:** yes (more than one)

## Setup

```bash
./cka/assignments/20-broken-service/setup.sh
```

Don't read `setup.sh` before solving: it shows the answer.

## Task

`shop/frontend` pods are `Running`, but `http://frontend.shop` doesn't answer from inside the cluster.

1. Find every root cause. Use `kubectl get endpointslices`, `kubectl describe` and a client pod.
2. Fix the Service only. Don't touch or recreate the Deployment.
3. Write the root causes to `/opt/cka/svc-fix.txt`.

## Done when

- `kubectl -n shop get endpointslices -l kubernetes.io/service-name=frontend` lists both pod IPs.
- `wget -qO- http://frontend.shop` from a busybox pod returns the nginx welcome page.

## Solution

In `solution/`, write each root cause, the command that revealed it, and the fix.
