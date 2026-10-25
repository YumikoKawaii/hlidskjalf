# 06 · CRDs & operators

**Domain:** Cluster Architecture, Installation & Configuration
**Planted faults:** none

## Setup

```bash
./cka/assignments/06-crds-operators/setup.sh
```

This installs the latest cert-manager release (its CRDs and controllers).

## Task

1. Write the names of every CRD in the `cert-manager.io` group to `/opt/cka/cm-crds.txt`, one per line.
2. Write the output of `kubectl explain` for `Certificate.spec` to `/opt/cka/cert-spec.txt`.
3. Create a ClusterIssuer `selfsigned` that uses the self-signed issuer type.
4. Create a Certificate `web-tls` in `default` for DNS name `web.lab.local`, issued by `selfsigned`, stored in Secret `web-tls`.
5. Write the certificate's expiry time, read from the Certificate's status, to `/opt/cka/web-tls-expiry.txt`.

## Done when

- `/opt/cka/cm-crds.txt` lists exactly the `cert-manager.io` CRDs.
- `kubectl get certificate web-tls` shows `READY=True`.
- `kubectl get secret web-tls -o jsonpath='{.type}'` returns `kubernetes.io/tls`.

## Solution

Put your manifests and commands in `solution/`.
