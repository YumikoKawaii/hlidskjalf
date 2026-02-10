#!/bin/bash

ACTION=${1:-down}
NAMESPACE=observe

case "$ACTION" in
  down)
    echo "Scaling down observe stack..."
    kubectl scale statefulset prometheus loki grafana --replicas=0 -n "$NAMESPACE"
    kubectl scale deployment otel-collector --replicas=0 -n "$NAMESPACE"
    kubectl patch daemonset promtail -n "$NAMESPACE" -p '{"spec":{"template":{"spec":{"nodeSelector":{"non-existing":"true"}}}}}'
    echo "Done. All observe pods scaling to zero."
    ;;
  up)
    echo "Scaling up observe stack..."
    kubectl scale statefulset prometheus loki grafana --replicas=1 -n "$NAMESPACE"
    kubectl scale deployment otel-collector --replicas=1 -n "$NAMESPACE"
    kubectl patch daemonset promtail -n "$NAMESPACE" --type=json -p='[{"op":"remove","path":"/spec/template/spec/nodeSelector"}]' 2>/dev/null
    echo "Done. All observe pods scaling up."
    ;;
  *)
    echo "Usage: $0 [up|down]"
    exit 1
    ;;
esac
