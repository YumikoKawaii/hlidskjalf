# 12 · Services & DNS

**Domain:** Services & Networking
**Planted faults:** none

## Setup

```bash
./cka/assignments/12-services-dns/setup.sh
```

This creates `net/web` (2 replicas of nginx) and no Services.

## Task

1. Expose `web` as NodePort Service `web-np` on port 80 with node port `30080`. From `cp1`, curl every node's IP on `30080`.
2. Create a ClusterIP Service `web-int` that listens on `8080` and forwards to the pods' port 80.
3. From a busybox pod in a different namespace, fetch `http://web-int.net.svc.cluster.local:8080`. Save the `nslookup` output for that name to `/opt/cka/web-int-dns.txt`.
4. Create a headless Service `web-headless` for the same pods. Save the A records its name returns to `/opt/cka/headless.txt`.
5. Create an ExternalName Service `docs` in `net` that points to `kubernetes.io`.

## Done when

- `curl <any-node-ip>:30080` returns the nginx welcome page.
- `/opt/cka/web-int-dns.txt` shows the ClusterIP, and `/opt/cka/headless.txt` shows the two pod IPs.
- `nslookup docs.net.svc.cluster.local` returns a CNAME to `kubernetes.io`.

## Solution

Put your manifests and commands in `solution/`.
