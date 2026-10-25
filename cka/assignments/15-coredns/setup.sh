#!/usr/bin/env bash
source "$(dirname "$0")/../lib.sh"
preflight

on cp1 <<'EOF'
set -euo pipefail
img=$(kubectl -n kube-system get deploy coredns -o jsonpath='{.spec.template.spec.containers[0].image}')
kubectl delete ns corp-dns --ignore-not-found --wait=true
rm -f /opt/cka/Corefile
kubectl apply -f - <<YAML
apiVersion: v1
kind: Namespace
metadata:
  name: corp-dns
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: corp-dns
  namespace: corp-dns
data:
  Corefile: |
    corp.local:5353 {
        hosts {
            10.10.10.10 db.corp.local
            10.10.10.11 vault.corp.local
        }
        log
        errors
    }
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: corp-dns
  namespace: corp-dns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: corp-dns
  template:
    metadata:
      labels:
        app: corp-dns
    spec:
      containers:
        - name: coredns
          image: $img
          args: ["-conf", "/etc/coredns/Corefile"]
          ports:
            - name: dns
              containerPort: 5353
              protocol: UDP
            - name: dns-tcp
              containerPort: 5353
              protocol: TCP
          volumeMounts:
            - name: conf
              mountPath: /etc/coredns
      volumes:
        - name: conf
          configMap:
            name: corp-dns
---
apiVersion: v1
kind: Service
metadata:
  name: corp-dns
  namespace: corp-dns
spec:
  clusterIP: 10.96.100.53
  selector:
    app: corp-dns
  ports:
    - name: dns
      port: 53
      targetPort: 5353
      protocol: UDP
    - name: dns-tcp
      port: 53
      targetPort: 5353
      protocol: TCP
YAML
kubectl -n corp-dns rollout status deploy/corp-dns --timeout=120s
EOF
