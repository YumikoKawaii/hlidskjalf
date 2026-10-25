#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on_root worker1 <<'EOF'
set -euo pipefail
f=/var/lib/kubelet/config.yaml
if ! grep -q 'clientCAFile: /etc/kubernetes/pki/ca.crt' "$f"; then
  echo "kubelet config does not have the expected clientCAFile; restore it before re-running" >&2
  exit 1
fi
c=$(cat "$f")
printf '%s\n' "${c//pki\/ca.crt/pki\/ca-bundle.crt}" > "$f"
systemctl disable --now kubelet
EOF
