kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
## failed due to identity check between metrics server to kubelet on every node
## add serverTLSBootstrap: true to kubelet-config and config.yaml of kubelet on every nodes
## restart kubelet on every nodes
## get csr and approve those requests via kubectl get csr and kubectl certificate approve <cert-name>