#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
for ns in db backend frontend; do
  kubectl delete ns "$ns" --ignore-not-found --wait=true
done
rm -f /opt/cka/netpol.txt
kubectl apply -f - <<'YAML'
apiVersion: v1
kind: Namespace
metadata:
  name: db
---
apiVersion: v1
kind: Namespace
metadata:
  name: backend
  labels:
    tier: backend
---
apiVersion: v1
kind: Namespace
metadata:
  name: frontend
  labels:
    tier: frontend
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-stub
  namespace: db
data:
  default.conf: |
    server {
      listen 5432;
      location / {
        default_type text/plain;
        return 200 "postgres ok\n";
      }
    }
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres
  namespace: db
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
        - name: postgres
          image: nginx:1.27
          ports:
            - containerPort: 5432
          volumeMounts:
            - name: conf
              mountPath: /etc/nginx/conf.d
      volumes:
        - name: conf
          configMap:
            name: postgres-stub
---
apiVersion: v1
kind: Service
metadata:
  name: postgres
  namespace: db
spec:
  selector:
    app: postgres
  ports:
    - port: 5432
---
apiVersion: v1
kind: Pod
metadata:
  name: api
  namespace: backend
  labels:
    role: api
spec:
  containers:
    - name: shell
      image: busybox:1.37
      command: ["sleep", "infinity"]
---
apiVersion: v1
kind: Pod
metadata:
  name: worker
  namespace: backend
  labels:
    role: worker
spec:
  containers:
    - name: shell
      image: busybox:1.37
      command: ["sleep", "infinity"]
---
apiVersion: v1
kind: Pod
metadata:
  name: api
  namespace: frontend
  labels:
    role: api
spec:
  containers:
    - name: shell
      image: busybox:1.37
      command: ["sleep", "infinity"]
YAML
EOF
