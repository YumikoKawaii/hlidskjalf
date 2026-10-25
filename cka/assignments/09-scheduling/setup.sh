#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on_root worker1 <<'EOF'
set -euo pipefail
rm -f /etc/kubernetes/manifests/static-web.yaml
EOF

node=$(node_name worker2)
{ echo "node=$node"; cat <<'EOF'
set -euo pipefail
kubectl delete ns sched --ignore-not-found --wait=true
kubectl create ns sched
kubectl label node "$node" accel- >/dev/null 2>&1 || true
kubectl taint node "$node" dedicated- >/dev/null 2>&1 || true
EOF
} | on cp1
