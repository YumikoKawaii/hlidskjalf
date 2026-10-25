#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
if ! command -v helm >/dev/null; then
  ver=$(curl -fsSL https://get.helm.sh/helm-latest-version)
  arch=$(dpkg --print-architecture)
  tmp=$(mktemp -d)
  curl -fsSL "https://get.helm.sh/helm-${ver}-linux-${arch}.tar.gz" | tar xz -C "$tmp"
  sudo install -m 0755 "$tmp/linux-${arch}/helm" /usr/local/bin/helm
  rm -rf "$tmp"
fi
helm -n web-apps uninstall edge >/dev/null 2>&1 || true
helm repo remove podinfo >/dev/null 2>&1 || true
kubectl delete ns web-apps --ignore-not-found --wait=true
rm -f /opt/cka/edge.yaml /opt/cka/edge-values.yaml
EOF
