# 11 · ResourceQuota & LimitRange

**Domain:** Workloads & Scheduling
**Planted faults:** none

## Setup

```bash
./cka/assignments/11-quota-limitrange/setup.sh
```

This creates `team-a/hog` (2 replicas, no resource requests or limits).

## Task

1. Create a ResourceQuota `team-a-quota`: `requests.cpu=2`, `requests.memory=2Gi`, `limits.cpu=4`, `limits.memory=4Gi`, `pods=10`.
2. Create a LimitRange `team-a-limits` for containers: default request `100m`/`128Mi`, default limit `200m`/`256Mi`, max `1` CPU and `1Gi`.
3. Make the existing `hog` pods pick up the defaults without editing the Deployment.
4. Scale `hog` to 15. Write which quota line stops it, and where you saw the error, to `/opt/cka/quota.txt`.
5. Try to create a pod in `team-a` that requests 2 CPUs. Record the rejection message in the same file.

## Done when

- `kubectl -n team-a describe quota team-a-quota` shows `pods 10/10`.
- Every `hog` pod has requests and limits set.
- `/opt/cka/quota.txt` names the blocking limit and the LimitRange rejection.

## Solution

Put your manifests and commands in `solution/`.
