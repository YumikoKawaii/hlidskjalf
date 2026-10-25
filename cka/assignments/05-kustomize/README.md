# 05 · Kustomize

**Domain:** Cluster Architecture, Installation & Configuration
**Planted faults:** none

## Setup

```bash
./cka/assignments/05-kustomize/setup.sh
```

This creates `/opt/cka/kustomize/app.properties` on `cp1`.

## Task

Work in `/opt/cka/kustomize` on `cp1`.

1. `base/`: a Deployment `api` (`nginx:1.27`, 1 replica) and a ClusterIP Service `api` on port 80.
2. `overlays/prod/`, built on the base:
   - namespace `prod` (the overlay creates it)
   - name prefix `prod-`
   - label `env=prod` on every resource and in selectors
   - 3 replicas
   - image swapped to `nginx:1.27-alpine` with the `images` field, not a patch
   - a ConfigMap `app-config` generated from `app.properties`, mounted into the container at `/etc/app`. The generator source has to sit inside the overlay directory, so copy the file there.
3. Apply it with `kubectl apply -k`.
4. Change `log.level` to `debug`, apply again, and watch what happens to the ConfigMap name and the pods.

## Done when

- `kubectl -n prod get deploy prod-api` shows 3/3 on `nginx:1.27-alpine`.
- `kubectl -n prod get cm` shows two hashed `prod-app-config-*` ConfigMaps, and the pods mount the newer one.
- `kubectl -n prod exec deploy/prod-api -- cat /etc/app/app.properties` shows `log.level=debug`.

## Solution

Copy the `base/` and `overlays/` trees into `solution/` and add notes on step 4.
