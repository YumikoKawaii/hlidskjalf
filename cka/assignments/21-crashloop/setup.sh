#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns crash --ignore-not-found --wait=true
rm -f /opt/cka/crash.txt
kubectl apply -f - <<'YAML'
apiVersion: v1
kind: Namespace
metadata:
  name: crash
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: payments-config
  namespace: crash
data:
  DB_URL: postgres://db.crash.svc:5432/payments
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: payments
  namespace: crash
spec:
  replicas: 1
  selector:
    matchLabels:
      app: payments
  template:
    metadata:
      labels:
        app: payments
    spec:
      containers:
        - name: app
          image: busybox:1.37
          command:
            - sh
            - -c
            - |
              if [ -z "$DB_URL" ]; then echo "FATAL: DB_URL is not set" >&2; exit 1; fi
              echo "connected to $DB_URL"
              sleep 3600
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: reports
  namespace: crash
spec:
  replicas: 1
  selector:
    matchLabels:
      app: reports
  template:
    metadata:
      labels:
        app: reports
    spec:
      containers:
        - name: app
          image: nginx:1.27
          livenessProbe:
            httpGet:
              path: /
              port: 8080
            initialDelaySeconds: 5
            periodSeconds: 5
            failureThreshold: 2
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: importer
  namespace: crash
spec:
  replicas: 1
  selector:
    matchLabels:
      app: importer
  template:
    metadata:
      labels:
        app: importer
    spec:
      containers:
        - name: app
          image: busybox:1.37
          command:
            - sh
            - -c
            - |
              echo "loading dataset"
              dd if=/dev/zero bs=1M count=200 2>/dev/null | tail > /dev/null
              echo "dataset loaded"
              sleep 3600
          resources:
            requests:
              memory: 32Mi
            limits:
              memory: 32Mi
YAML
EOF
