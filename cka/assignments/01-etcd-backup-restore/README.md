# 01 · etcd backup & restore

**Domain:** Cluster Architecture, Installation & Configuration
**Planted faults:** none

## Setup

```bash
./cka/assignments/01-etcd-backup-restore/setup.sh
```

The script installs `etcdctl` and `etcdutl` on `cp1` (matching the running etcd version) if they're missing, and clears any earlier attempt.

## Task

On `cp1`:

1. Take a snapshot of the running etcd and save it to `/opt/backup/etcd-snap.db`. Use the certificates under `/etc/kubernetes/pki/etcd`.
2. Check the snapshot's status (hash, revision, total keys).
3. Create a namespace `canary`.
4. Restore the snapshot into `/var/lib/etcd-restore`.
5. Point the etcd static pod at the restored data directory and wait for the control plane to come back.

## Done when

- `kubectl get ns canary` returns `NotFound`.
- `kubectl get nodes` shows all three nodes `Ready`.
- `/etc/kubernetes/manifests/etcd.yaml` mounts `/var/lib/etcd-restore`.

## Solution

Put your commands and notes in `solution/`.
