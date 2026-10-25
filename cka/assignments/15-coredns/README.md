# 15 · CoreDNS

**Domain:** Services & Networking
**Planted faults:** none

## Setup

```bash
./cka/assignments/15-coredns/setup.sh
```

This deploys a "corporate" DNS server at `10.96.100.53` that's authoritative for `corp.local` (`db.corp.local` → `10.10.10.10`, `vault.corp.local` → `10.10.10.11`). Cluster DNS doesn't know about it yet.

## Task

1. Configure the cluster CoreDNS so that `corp.local` queries go to `10.96.100.53`. Pod specs stay untouched.
2. Make `legacy.lab` resolve to `10.20.30.40` from every pod, served by the cluster CoreDNS.
3. Create a pod `isolated` in `default` with `dnsPolicy: None` that uses only `10.96.100.53` as its nameserver. Show that it resolves `db.corp.local` but not `kubernetes.default`.
4. Save the final cluster Corefile to `/opt/cka/Corefile`.

## Done when

- `nslookup db.corp.local` from any regular pod returns `10.10.10.10`.
- `nslookup legacy.lab` from any regular pod returns `10.20.30.40`.
- `kubernetes.default.svc.cluster.local` still resolves from regular pods.
- The CoreDNS pods are healthy, with no errors about the Corefile in their logs.

## Solution

Put your ConfigMap changes and the `isolated` pod manifest in `solution/`.
