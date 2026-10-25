#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns ci --ignore-not-found --wait=true
kubectl delete clusterrolebinding deployer-node-viewer --ignore-not-found
kubectl delete clusterrole node-viewer --ignore-not-found
kubectl create ns ci
kubectl -n ci create deployment builder --image=nginx:1.27
kubectl -n ci create secret generic registry-token --from-literal=token=not-a-real-token
rm -f /opt/cka/deployer.kubeconfig
EOF
