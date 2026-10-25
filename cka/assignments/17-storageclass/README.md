# 17 · StorageClass

**Domain:** Storage
**Planted faults:** none

## Setup

```bash
./cka/assignments/17-storageclass/setup.sh
```

This installs Rancher's local-path-provisioner (provisioner `rancher.io/local-path`), creates namespace `sc`, and adds a StorageClass `slow` that is marked as default.

## Task

1. Create a StorageClass `fast`: provisioner `rancher.io/local-path`, `volumeBindingMode: WaitForFirstConsumer`, `reclaimPolicy: Retain`.
2. Make `fast` the cluster's only default StorageClass.
3. In `sc`, create a PVC `data` (1Gi, `ReadWriteOnce`) **without** setting `storageClassName`. Check its status before any pod uses it, and explain what you see.
4. Create a pod `writer` (`busybox:1.37`) that mounts `data` at `/data` and writes `/data/hello.txt`.
5. Try to expand `data` to 2Gi. Save the error to `/opt/cka/expand.txt` and note what would have to change for expansion to work.
6. Delete `writer` and `data`. Check what happens to the PV.

## Done when

- `kubectl get sc` shows `(default)` only on `fast`.
- `data` was `Pending` until `writer` started, then `Bound` to a PV of class `fast`.
- After step 6, the PV still exists with status `Released`.

## Solution

Put your manifests in `solution/`, plus notes on steps 3 and 5.
