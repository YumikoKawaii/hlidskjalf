# 19 · Broken scheduler

**Domain:** Troubleshooting
**Planted faults:** yes (more than one, layered)

## Setup

```bash
./cka/assignments/19-broken-scheduler/setup.sh
```

Don't read `setup.sh` before solving: it shows the answer.

## Task

New pods in `sched-down/pending-app` stay `Pending`. Fix the cluster so they schedule. The second fault only shows up once the first is fixed.

1. Find every root cause.
2. Fix them, and clean up anything the broken state left behind on the node.

## Done when

- `kubectl -n sched-down get pods` shows 2/2 `Running`.
- `kubectl -n kube-system get pod -l component=kube-scheduler` shows the scheduler `Running` with no restarts climbing.
- `/etc/kubernetes/manifests/kube-scheduler.yaml` matches what kubeadm would have generated.

## Solution

In `solution/`, write each root cause, the command that revealed it, and the fix.
