kubectl create sa deployer -n ci
kubectl create role deployer -n ci --verb=get,list,create,update,patch --resource=deployment,statefulset,deployment/scale
kubectl create rolebinding -n ci deployer-binding --role=deployer --serviceaccount=ci:deployer
kubectl create clusterrole node-viewer --verb=get,list --resource=node
kubectl create clusterrolebinding deployer-node-viewer --clusterrole=node-viewer --serviceaccount=ci:deployer
kubectl create token deployer -n ci --duration=24h > token.txt
CLUSTER_NAME=$(kubectl config view --minify -o jsonpath="{.clusters[0].name}")
SERVER=$(kubectl config view --minify -o jsonpath="{.clusters[0].cluster.server}")
TARGET=/opt/cka/deployer.kubeconfig
kubectl config view --minify --raw -o jsonpath="{.clusters[0].cluster.certificate-authority-data}" | base64 -d > ca.crt
kubectl config set-cluster "$CLUSTER_NAME" --server="$SERVER" --certificate-authority=ca.crt --embed-certs=true --kubeconfig="$TARGET"
kubectl config set-credentials deployer --token=$(cat token.txt) --kubeconfig="$TARGET"
kubectl config set-context deployer@"$CLUSTER_NAME" --cluster="$CLUSTER_NAME" --user=deployer --namespace=ci --kubeconfig="$TARGET"
rm ca.crt
rm token.txt
export KUBECONFIG=$TARGET
kubectl config use-context deployer@kubernetes
kubectl -n ci scale deploy builder --replicas=2