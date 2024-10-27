# set namespace to elysium
kubectl config set-context --current --namespace=elysium

# deploy mysql
helm install mysql ./infrastructure/mysql

# deploy redis
helm install redis ./infrastructure/redis
