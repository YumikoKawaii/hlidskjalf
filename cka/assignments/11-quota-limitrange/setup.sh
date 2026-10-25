#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns team-a --ignore-not-found --wait=true
kubectl create ns team-a
kubectl -n team-a create deployment hog --image=nginx:1.27 --replicas=2
rm -f /opt/cka/quota.txt
EOF
