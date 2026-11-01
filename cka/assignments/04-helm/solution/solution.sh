helm repo add podinfo https://stefanprodan.github.io/podinfo
helm install edge podinfo/podinfo -n web-apps --create-namespace --version 6.16.0 --set replicaCount=2
helm upgrade edge podinfo/podinfo --version 6.17.0 -n web-apps
helm rollback edge -n web-apps 1
helm get values edge -n web-apps -o yaml > /opt/cka/edge-values.yaml
helm template edge podinfo/podinfo -n web-apps --version 6.17.0 --set replicaCount=3 > /opt/cka/edge.yaml