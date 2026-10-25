set -euo pipefail

on() { limactl shell --workdir / "$1" bash -s; }

on_root() { limactl shell --workdir / "$1" sudo bash -s; }

node_name() { limactl shell --workdir / "$1" hostname; }

preflight() {
  on cp1 <<'EOF'
set -euo pipefail
kubectl get nodes >/dev/null
sudo install -d -o "$(id -un)" -g "$(id -gn)" /opt/cka
EOF
}
