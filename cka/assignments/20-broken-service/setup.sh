#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns shop --ignore-not-found --wait=true
rm -f /opt/cka/svc-fix.txt
kubectl apply -f - <<'YAML'
apiVersion: v1
kind: Namespace
metadata:
  name: shop
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  namespace: shop
spec:
  replicas: 2
  selector:
    matchLabels:
      app: frontend
      tier: web
  template:
    metadata:
      labels:
        app: frontend
        tier: web
    spec:
      containers:
        - name: nginx
          image: nginx:1.27
          ports:
            - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: frontend
  namespace: shop
spec:
  selector:
    app: front-end
    tier: web
  ports:
    - port: 80
      targetPort: 8080
YAML
EOF
