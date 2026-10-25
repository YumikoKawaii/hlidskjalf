# 18 · Broken kubelet

**Domain:** Troubleshooting
**Planted faults:** yes (more than one)

## Setup

```bash
./cka/assignments/18-broken-kubelet/setup.sh
```

Don't read `setup.sh` before solving: it shows the answer.

## Task

The node behind VM `worker1` went `NotReady`. Get it back without rebuilding or rejoining it.

1. Find every reason the node is down.
2. Fix them so the node is `Ready` and stays that way after a reboot.

## Done when

- `kubectl get nodes` shows the `worker1` node `Ready`.
- `systemctl is-enabled kubelet` on `worker1` returns `enabled`.
- After `limactl stop worker1 && limactl start worker1`, the node comes back `Ready` by itself.

## Solution

In `solution/`, write each root cause, the command that revealed it, and the fix.
