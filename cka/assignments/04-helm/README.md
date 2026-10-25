# 04 · Helm

**Domain:** Cluster Architecture, Installation & Configuration
**Planted faults:** none

## Setup

```bash
./cka/assignments/04-helm/setup.sh
```

This installs `helm` on `cp1` if it's missing and removes anything left from an earlier attempt.

## Task

1. Add the chart repo `podinfo` at `https://stefanprodan.github.io/podinfo`.
2. Install the chart `podinfo/podinfo` as release `edge` in namespace `web-apps` (create the namespace through Helm). Pin the **second-newest** chart version and set `replicaCount=2`.
3. Upgrade `edge` to the newest chart version, keeping `replicaCount=2`.
4. Roll `edge` back to revision 1.
5. Save the values the release currently runs with to `/opt/cka/edge-values.yaml`.
6. Render the newest chart version with `replicaCount=3` to `/opt/cka/edge.yaml` without installing anything.

## Done when

- `helm -n web-apps history edge` shows 3 revisions, and the latest is a rollback to 1.
- `helm -n web-apps list` shows the second-newest chart version.
- `kubectl -n web-apps get deploy` shows 2 replicas ready.
- `/opt/cka/edge.yaml` contains a Deployment with `replicas: 3`, and no third replica exists in the cluster.

## Solution

Put your commands in `solution/`.
