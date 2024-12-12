kubectl create configmap locust-script --from-file=./applications/locust/locustfile.py
kubectl apply -f ./applications/locust/master/deployment.yaml
kubectl apply -f ./applications/locust/master/service.yaml
kubectl apply -f ./applications/locust/workers/deployment.yaml