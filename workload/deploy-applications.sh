# deploy authenticator
kubectl apply -f ./applications/authenticator/deployment.yaml
kubectl apply -f ./applications/authenticator/service.yaml

# deploy posts
kubectl apply -f ./applications/posts/deployment.yaml
kubectl apply -f ./applications/posts/service.yaml

# deploy users
kubectl apply -f ./applications/users/deployment.yaml
kubectl apply -f ./applications/users/service.yaml

# deploy interactions
kubectl apply -f ./applications/interactions/deployment.yaml
kubectl apply -f ./applications/interactions/service.yaml

# deploy center