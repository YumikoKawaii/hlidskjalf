# 10 · Native sidecar

**Domain:** Workloads & Scheduling
**Planted faults:** none

## Setup

```bash
./cka/assignments/10-native-sidecar/setup.sh
```

This creates `sidecar/web`, an nginx container that writes its logs to an `emptyDir` at `/var/log/nginx`, and a Service `web`.

## Task

1. Add a sidecar `log-tailer` (`busybox:1.37`) to `web` that runs `tail -F /var/log/nginx/access.log` from the shared volume. Use native sidecar syntax: an `initContainers` entry with `restartPolicy: Always`.
2. Send a few requests to `http://web.sidecar` from another pod.
3. Show that the requests appear in the `log-tailer` logs.

## Done when

- `kubectl -n sidecar get pod -l app=web` shows `2/2` ready.
- `kubectl -n sidecar logs deploy/web -c log-tailer` shows your requests.

## Solution

Put the manifest in `solution/`, plus a note on how a native sidecar differs from a second regular container: start order, and what happens when the pod belongs to a Job.
