#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns autoscale --ignore-not-found --wait=true
kubectl create ns autoscale
kubectl -n autoscale create deployment web --image=nginx:1.27
kubectl -n autoscale set resources deployment web --requests=cpu=50m --limits=cpu=200m
kubectl -n autoscale expose deployment web --port=80
rm -f /opt/cka/hpa.txt
EOF
