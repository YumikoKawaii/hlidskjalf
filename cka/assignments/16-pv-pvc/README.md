# 16 · PV & PVC

**Domain:** Storage
**Planted faults:** none

## Setup

```bash
./cka/assignments/16-pv-pvc/setup.sh
```

This creates namespace `storage` and the directory `/mnt/log` on VM `worker1`, containing `seed.txt`.

## Task

1. Create a PersistentVolume `pv-log`: 1Gi, `ReadWriteOnce`, reclaim policy `Retain`, `storageClassName: manual`, a `local` volume at `/mnt/log`, with node affinity to the `worker1` node.
2. Create a PVC `pvc-log` in `storage` that binds to `pv-log`.
3. Create a pod `log-reader` (`busybox:1.37`) in `storage` that mounts the claim at `/log` and prints `/log/seed.txt` at startup.
4. Delete the pod and the PVC. Check the PV's phase.
5. Make `pv-log` `Available` again without deleting and recreating it.

## Done when

- `kubectl -n storage logs log-reader` printed `hello from <worker1 hostname>` (before step 4).
- After step 4, `kubectl get pv pv-log` shows `Released`, and after step 5 it shows `Available`.
- `seed.txt` still exists on `worker1`.

## Solution

Put your manifests and commands in `solution/`.
