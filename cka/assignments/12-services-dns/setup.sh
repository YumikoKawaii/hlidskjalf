#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns net --ignore-not-found --wait=true
kubectl create ns net
kubectl -n net create deployment web --image=nginx:1.27 --replicas=2
rm -f /opt/cka/web-int-dns.txt /opt/cka/headless.txt
EOF
