# 23 · Node network

**Domain:** Troubleshooting
**Planted faults:** yes

## Setup

```bash
./cka/assignments/23-node-network/setup.sh
```

Don't read `setup.sh` before solving: it shows the answer. Rebooting the VM also clears the fault, which doesn't count as a fix.

## Task

In `netcheck`, `probe-a` and `probe-b` are identical busybox pods on different nodes. `target` runs on the same node as `probe-b`.

- `kubectl -n netcheck exec probe-b -- wget -qO- -T 3 http://target` works.
- `kubectl -n netcheck exec probe-a -- wget -qO- -T 3 http://target` doesn't.

1. Work out which layer is broken (pod, Service, DNS, CNI, node) and on which node. Use Cilium's own status tools as well as kubectl.
2. Fix it on the node, without restarting Cilium or the VM.

## Done when

- Both probes reach `http://target` and `http://target.netcheck.svc.cluster.local`.
- `cilium status` (or `cilium-health status` inside a Cilium agent pod) reports every node reachable.

## Solution

In `solution/`, write how you narrowed it down layer by layer, the root cause and the fix.
