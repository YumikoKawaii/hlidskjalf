# 21 · CrashLoopBackOff

**Domain:** Troubleshooting
**Planted faults:** yes (one per Deployment, each a different kind)

## Setup

```bash
./cka/assignments/21-crashloop/setup.sh
```

Don't read `setup.sh` before solving: it shows the answer.

## Task

None of the three Deployments in `crash` (`payments`, `reports`, `importer`) stay up.

1. For each one, find why its container keeps restarting. `kubectl logs --previous`, `kubectl describe` and the container's last state are your tools.
2. Write one line per Deployment to `/opt/cka/crash.txt`: `name: cause (evidence)`.
3. Fix each one with the smallest change that addresses the cause:
   - `payments` must take its config from the existing ConfigMap `payments-config`.
   - `reports` must keep a liveness probe.
   - `importer` must keep a memory limit.

## Done when

- `kubectl -n crash get pods` shows all three `Running`, with restart counts that stay put for 2 minutes.
- `kubectl -n crash logs deploy/payments` prints `connected to postgres://…`.
- `kubectl -n crash logs deploy/importer` prints `dataset loaded`.

## Solution

Put your patches in `solution/`, plus how you told the three failure types apart.
