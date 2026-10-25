#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
cur=$(kubeadm version -o short)
minor=$(echo "$cur" | cut -d. -f2)
next="v1.$((minor + 1))"
if ! curl -fsIL "https://pkgs.k8s.io/core:/stable:/${next}/deb/Release" >/dev/null; then
  echo "cluster is on $cur and $next is not published; rebuild the nodes with an older minor in bootstrap.sh" >&2
  exit 1
fi
kubectl create ns upgrade --dry-run=client -o yaml | kubectl apply -f -
kubectl -n upgrade create deployment canary --image=nginx:1.27 --replicas=4 --dry-run=client -o yaml | kubectl apply -f -
echo "upgrade target: $next"
EOF
