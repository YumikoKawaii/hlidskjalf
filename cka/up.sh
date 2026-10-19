#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"
K8S_MINOR="${K8S_MINOR:-1.36}"

NODES=("$@")
[ ${#NODES[@]} -eq 0 ] && NODES=(cp1 worker1 worker2)

for name in "${NODES[@]}"; do
  case "$name" in
    cp1) mem=4 ;;
    *) mem=3 ;;
  esac
  if ! limactl list -q 2>/dev/null | grep -qx "$name"; then
    limactl create --tty=false --name="$name" --vm-type=vz \
      --cpus=2 --memory="$mem" --disk=20 \
      --network=lima:user-v2 --plain template:ubuntu-24.04
  fi
  if [ "$(limactl list "$name" --format '{{.Status}}')" != "Running" ]; then
    limactl start --tty=false "$name"
  fi
  limactl shell "$name" sudo bash -s -- "$K8S_MINOR" <node-setup.sh
done

limactl list
