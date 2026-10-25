#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

read -r mode proto port < <(on cp1 <<'EOF'
set -euo pipefail
cm() { kubectl -n kube-system get cm cilium-config -o go-template="{{or (index .data \"$1\") \"$2\"}}"; }
echo "$(cm routing-mode tunnel) $(cm tunnel-protocol vxlan) $(cm tunnel-port 0)"
EOF
)

if [ "$mode" != "tunnel" ]; then
  echo "cilium routing-mode is $mode; this assignment needs tunnel mode" >&2
  exit 1
fi
if [ "$port" = "0" ]; then
  if [ "$proto" = "geneve" ]; then port=6081; else port=8472; fi
fi

{ echo "port=$port"; cat <<'EOF'
set -euo pipefail
iptables -I INPUT -p udp --dport "$port" -j DROP
iptables -I OUTPUT -p udp --dport "$port" -j DROP
EOF
} | on_root worker1

w1=$(node_name worker1)
w2=$(node_name worker2)
{ echo "w1=$w1"; echo "w2=$w2"; cat <<'EOF'
set -euo pipefail
kubectl delete ns netcheck --ignore-not-found --wait=true
kubectl create ns netcheck
kubectl apply -f - <<YAML
apiVersion: v1
kind: Pod
metadata:
  name: probe-a
  namespace: netcheck
spec:
  nodeSelector:
    kubernetes.io/hostname: $w1
  containers:
    - name: shell
      image: busybox:1.37
      command: ["sleep", "infinity"]
---
apiVersion: v1
kind: Pod
metadata:
  name: probe-b
  namespace: netcheck
spec:
  nodeSelector:
    kubernetes.io/hostname: $w2
  containers:
    - name: shell
      image: busybox:1.37
      command: ["sleep", "infinity"]
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: target
  namespace: netcheck
spec:
  replicas: 1
  selector:
    matchLabels:
      app: target
  template:
    metadata:
      labels:
        app: target
    spec:
      nodeSelector:
        kubernetes.io/hostname: $w2
      containers:
        - name: nginx
          image: nginx:1.27
---
apiVersion: v1
kind: Service
metadata:
  name: target
  namespace: netcheck
spec:
  selector:
    app: target
  ports:
    - port: 80
YAML
EOF
} | on cp1
