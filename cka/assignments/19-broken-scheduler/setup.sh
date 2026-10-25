#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on_root cp1 <<'EOF'
set -euo pipefail
f=/etc/kubernetes/manifests/kube-scheduler.yaml
if ! grep -q 'image: registry.k8s.io/kube-scheduler:' "$f"; then
  echo "kube-scheduler manifest is not in its original state; restore it before re-running" >&2
  exit 1
fi
c=$(cat "$f")
c=${c//registry.k8s.io\/kube-scheduler:/registry.k8s.io\/kube-schedular:}
c=${c//scheduler.conf/scheduler.kubeconfig}
printf '%s\n' "$c" > /root/kube-scheduler.yaml
mv /root/kube-scheduler.yaml "$f"
EOF

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns sched-down --ignore-not-found --wait=true
kubectl create ns sched-down
kubectl -n sched-down create deployment pending-app --image=nginx:1.27 --replicas=2
EOF
