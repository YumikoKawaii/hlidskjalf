#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns sidecar --ignore-not-found --wait=true
kubectl create ns sidecar
kubectl apply -f - <<'YAML'
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web
  namespace: sidecar
spec:
  replicas: 1
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      containers:
        - name: nginx
          image: nginx:1.27
          volumeMounts:
            - name: logs
              mountPath: /var/log/nginx
      volumes:
        - name: logs
          emptyDir: {}
YAML
kubectl -n sidecar expose deployment web --port=80
EOF
