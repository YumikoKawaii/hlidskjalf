#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns storage --ignore-not-found --wait=true
kubectl delete pv pv-log --ignore-not-found
kubectl create ns storage
EOF

on_root worker1 <<'EOF'
set -euo pipefail
rm -rf /mnt/log
mkdir -p /mnt/log
echo "hello from $(hostname)" > /mnt/log/seed.txt
EOF
