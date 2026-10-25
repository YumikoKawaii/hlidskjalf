# 09 · Scheduling

**Domain:** Workloads & Scheduling
**Planted faults:** none

## Setup

```bash
./cka/assignments/09-scheduling/setup.sh
```

This creates namespace `sched` and removes earlier labels, taints and static pods.

## Task

1. Label the node behind VM `worker2` with `accel=gpu` and taint it `dedicated=gpu:NoSchedule`.
2. In `sched`, create a pod `gpu-job` (`busybox:1.37`, `sleep 3600`). It must run only on nodes labelled `accel=gpu` (required node affinity) and must tolerate the taint.
3. In `sched`, create a pod `anywhere` (`busybox:1.37`, `sleep 3600`) with a **preferred** node affinity for `accel=gpu` and no toleration. Note where it lands and why.
4. On VM `worker1`, run `static-web` (`nginx:1.27`) as a static pod.
5. Write each pod's node to `/opt/cka/placement.txt` as `pod node`.

## Done when

- `gpu-job` runs on the `worker2` node.
- `anywhere` runs on the `worker1` node.
- `kubectl get pods -A` shows `static-web-<worker1 node>`. Deleting it with kubectl doesn't remove it.

## Solution

Put your manifests in `solution/`, plus a note on why `anywhere` didn't follow its preference.
