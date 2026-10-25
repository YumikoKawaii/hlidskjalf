#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml
kubectl -n cert-manager rollout status deploy/cert-manager --timeout=180s
kubectl -n cert-manager rollout status deploy/cert-manager-webhook --timeout=180s
kubectl -n default delete certificate web-tls --ignore-not-found
kubectl -n default delete secret web-tls --ignore-not-found
kubectl delete clusterissuer selfsigned --ignore-not-found
rm -f /opt/cka/cm-crds.txt /opt/cka/cert-spec.txt /opt/cka/web-tls-expiry.txt
EOF
