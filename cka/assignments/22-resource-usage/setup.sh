#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
if ! kubectl top nodes >/dev/null 2>&1; then
  echo "metrics API is not available; finish assignment 08 first" >&2
  exit 1
fi
kubectl delete ns load-a load-b --ignore-not-found --wait=true
rm -f /opt/cka/top-pod.txt /opt/cka/top-mem.txt /opt/cka/top-node.txt /opt/cka/throttle.txt
kubectl apply -f - <<'YAML'
apiVersion: v1
kind: Namespace
metadata:
  name: load-a
---
apiVersion: v1
kind: Namespace
metadata:
  name: load-b
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cpu-hog
  namespace: load-a
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cpu-hog
  template:
    metadata:
      labels:
        app: cpu-hog
    spec:
      containers:
        - name: app
          image: busybox:1.37
          command: ["sh", "-c", "while :; do :; done"]
          resources:
            requests:
              cpu: 100m
            limits:
              cpu: 400m
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: idle
  namespace: load-a
spec:
  replicas: 2
  selector:
    matchLabels:
      app: idle
  template:
    metadata:
      labels:
        app: idle
    spec:
      containers:
        - name: app
          image: nginx:1.27
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mem-hog
  namespace: load-b
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mem-hog
  template:
    metadata:
      labels:
        app: mem-hog
    spec:
      containers:
        - name: app
          image: busybox:1.37
          command: ["sh", "-c", "x=$(dd if=/dev/zero bs=1M count=120 2>/dev/null | tr '\\000' a); sleep 3600"]
          resources:
            limits:
              memory: 384Mi
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cpu-lite
  namespace: load-b
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cpu-lite
  template:
    metadata:
      labels:
        app: cpu-lite
    spec:
      containers:
        - name: app
          image: busybox:1.37
          command: ["sh", "-c", "while :; do :; done"]
          resources:
            limits:
              cpu: 100m
YAML
EOF
