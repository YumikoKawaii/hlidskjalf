#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns prod --ignore-not-found --wait=true
rm -rf /opt/cka/kustomize
mkdir -p /opt/cka/kustomize
cat > /opt/cka/kustomize/app.properties <<'PROPS'
log.level=info
feature.checkout=true
PROPS
EOF
