# 02 · Cluster upgrade

**Domain:** Cluster Architecture, Installation & Configuration
**Planted faults:** none

## Setup

```bash
./cka/assignments/02-cluster-upgrade/setup.sh
```

The script checks that pkgs.k8s.io publishes the next minor after the one the cluster runs, then prints the upgrade target. If the next minor isn't published, it stops: rebuild the nodes with an older minor in `bootstrap.sh` first. It also deploys `upgrade/canary` (4 replicas) as a workload that has to keep running.

## Task

1. Upgrade the control-plane node `cp1` to the latest patch of the next minor with `kubeadm`.
2. Upgrade `kubelet` and `kubectl` on `cp1` to the same version. Drain the node before the kubelet upgrade and uncordon it afterwards.
3. Leave `worker1` and `worker2` on the old version.
4. Keep the packages held (`apt-mark`) after the upgrade.

## Done when

- `kubectl get nodes` shows `cp1` on the new version, Ready and schedulable, and both workers on the old version.
- `kubectl version` reports the new server version.
- `kubectl -n upgrade get deploy canary` shows 4/4 available.
- `apt-mark showhold` on `cp1` lists `kubeadm`, `kubelet` and `kubectl`.

## Solution

Put your commands and notes in `solution/`.
