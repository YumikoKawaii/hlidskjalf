# 13 · NetworkPolicy

**Domain:** Services & Networking
**Planted faults:** none

## Setup

```bash
./cka/assignments/13-network-policy/setup.sh
```

This creates:

- `db/postgres`: a stub that answers HTTP on port 5432 (`postgres ok`), behind Service `postgres`
- `backend` (labelled `tier=backend`): pods `api` (`role=api`) and `worker` (`role=worker`)
- `frontend` (labelled `tier=frontend`): pod `api` (`role=api`)

Test with `kubectl -n <ns> exec <pod> -- wget -qO- -T 2 http://postgres.db.svc.cluster.local:5432`.

## Task

1. In `db`, create a NetworkPolicy `default-deny` that blocks all ingress and egress for every pod.
2. In `db`, create a NetworkPolicy `postgres-ingress` that allows traffic to `app=postgres` on TCP 5432 only from pods labelled `role=api` in namespaces labelled `tier=backend`.
3. In `db`, create a NetworkPolicy `allow-dns` that allows egress to CoreDNS only (UDP and TCP 53).
4. Run the test from all three client pods and write the results to `/opt/cka/netpol.txt` as `namespace/pod allowed|denied`.

## Done when

| From | Result |
|------|--------|
| `backend/api` | `postgres ok` |
| `backend/worker` | timeout |
| `frontend/api` | timeout |

## Solution

Put your manifests in `solution/`, plus a note on the difference between one `from` entry holding both selectors and two separate entries.
