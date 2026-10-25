#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns rollout --ignore-not-found --wait=true
kubectl create ns rollout
rm -f /opt/cka/rollout-history.txt
EOF
