kubectl rollout undo deployments/web --to-version=2
kubectl rollout history deployments/web > /opt/cka/rollout-history.txt