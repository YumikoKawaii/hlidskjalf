# 07 · Rollout & rollback

**Domain:** Workloads & Scheduling
**Planted faults:** none

## Setup

```bash
./cka/assignments/07-rollout-rollback/setup.sh
```

## Task

All in namespace `rollout`:

1. Create a Deployment `web` (`nginx:1.25`, 3 replicas) with a rolling update strategy of `maxSurge: 1` and `maxUnavailable: 0`.
2. Update the image to `nginx:1.27`. Record the reason in the `kubernetes.io/change-cause` annotation.
3. Update the image to `nginx:9.99` (it doesn't exist). Watch the rollout get stuck, and confirm the old pods keep serving.
4. Roll back to the revision that runs `nginx:1.27`. That's not revision 1.
5. Save `kubectl rollout history` for `web` to `/opt/cka/rollout-history.txt`.

## Done when

- `kubectl -n rollout get deploy web` shows 3/3 on `nginx:1.27`.
- The rollout history shows change causes for the 1.27 revision.
- At no point were fewer than 3 pods available.

## Solution

Put your commands in `solution/`, plus a note on why the broken rollout never took capacity down.
