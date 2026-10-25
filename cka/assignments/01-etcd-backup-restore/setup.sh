#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
if sudo grep -q etcd-restore /etc/kubernetes/manifests/etcd.yaml; then
  echo "etcd already runs from /var/lib/etcd-restore; point it back to /var/lib/etcd before re-running" >&2
  exit 1
fi
kubectl delete ns canary --ignore-not-found --wait=true
if ! command -v etcdctl >/dev/null || ! command -v etcdutl >/dev/null; then
  img=$(kubectl -n kube-system get pod -l component=etcd -o jsonpath='{.items[0].spec.containers[0].image}')
  tag=${img##*:}
  ver=${tag%%-*}
  name=etcd-v${ver}-linux-$(dpkg --print-architecture)
  tmp=$(mktemp -d)
  curl -fsSL "https://github.com/etcd-io/etcd/releases/download/v${ver}/${name}.tar.gz" | tar xz -C "$tmp"
  sudo install -m 0755 "$tmp/$name/etcdctl" "$tmp/$name/etcdutl" /usr/local/bin/
  rm -rf "$tmp"
fi
sudo install -d /opt/backup
sudo rm -f /opt/backup/etcd-snap.db
sudo rm -rf /var/lib/etcd-restore
EOF
