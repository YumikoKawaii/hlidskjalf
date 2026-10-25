#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
kubectl delete ns gw --ignore-not-found --wait=true
kubectl create ns gw
rm -f /opt/cka/canary.txt
for v in v1 v2; do
kubectl apply -f - <<YAML
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-$v
  namespace: gw
data:
  default.conf: |
    server {
      listen 80;
      location / {
        default_type text/plain;
        return 200 "$v\n";
      }
    }
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-$v
  namespace: gw
spec:
  replicas: 1
  selector:
    matchLabels:
      app: app-$v
  template:
    metadata:
      labels:
        app: app-$v
    spec:
      containers:
        - name: nginx
          image: nginx:1.27
          volumeMounts:
            - name: conf
              mountPath: /etc/nginx/conf.d
      volumes:
        - name: conf
          configMap:
            name: app-$v
---
apiVersion: v1
kind: Service
metadata:
  name: app-$v
  namespace: gw
spec:
  selector:
    app: app-$v
  ports:
    - port: 80
YAML
done
EOF
