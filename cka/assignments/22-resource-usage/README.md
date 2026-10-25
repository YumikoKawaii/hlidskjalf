# 22 · Resource usage

**Domain:** Troubleshooting
**Planted faults:** none

## Setup

```bash
./cka/assignments/22-resource-usage/setup.sh
```

This needs metrics-server from assignment 08. It starts several workloads in `load-a` and `load-b`. Give metrics a minute to show up.

## Task

1. Find the pod using the most CPU across all namespaces. Write it to `/opt/cka/top-pod.txt` as `namespace/name`.
2. Find the pod using the most memory in `load-b`. Write it to `/opt/cka/top-mem.txt`.
3. Find the node with the highest CPU percentage. Write its name to `/opt/cka/top-node.txt`.
4. Compare `cpu-hog`'s CPU usage with its spec, and write to `/opt/cka/throttle.txt` why it never goes past a certain value.

Use one command per answer where you can (`--sort-by`, `--no-headers`, `head`). Don't copy values by hand.

## Done when

- Each file holds a single line that matches what `kubectl top` shows.
- `/opt/cka/throttle.txt` names the setting that caps `cpu-hog`.

## Solution

Put the exact commands you used in `solution/`.
