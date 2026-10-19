#!/usr/bin/env bash
set -euo pipefail

K8S_MINOR="${1:?usage: node-setup.sh <k8s-minor e.g. 1.36>}"
export DEBIAN_FRONTEND=noninteractive

swapoff -a

cat >/etc/modules-load.d/k8s.conf <<EOF
overlay
br_netfilter
EOF
modprobe overlay
modprobe br_netfilter

cat >/etc/sysctl.d/k8s.conf <<EOF
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward = 1
EOF
sysctl --system >/dev/null

apt-get update -q
apt-get install -yq apt-transport-https ca-certificates curl gpg containerd etcd-client vim bash-completion jq

mkdir -p /etc/containerd
cat >/etc/containerd/config.toml <<EOF
version = 2
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
  runtime_type = "io.containerd.runc.v2"
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
  SystemdCgroup = true
EOF
systemctl restart containerd
systemctl enable containerd

mkdir -p /etc/apt/keyrings
curl -fsSL "https://pkgs.k8s.io/core:/stable:/v${K8S_MINOR}/deb/Release.key" | gpg --dearmor --yes -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v${K8S_MINOR}/deb/ /" >/etc/apt/sources.list.d/kubernetes.list
apt-get update -q
apt-get install -yq kubelet kubeadm kubectl cri-tools
apt-mark hold kubelet kubeadm kubectl
systemctl enable kubelet

cat >/etc/crictl.yaml <<EOF
runtime-endpoint: unix:///run/containerd/containerd.sock
image-endpoint: unix:///run/containerd/containerd.sock
EOF
