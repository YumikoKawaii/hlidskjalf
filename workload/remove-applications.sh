# remove authenticator
kubectl delete deployment authenticator
kubectl delete service authenticator

# remove posts
kubectl delete deployment posts
kubectl delete service posts

# remove users
kubectl delete deployment users
kubectl delete service users

# remove interactions
kubectl delete deployment interactions
kubectl delete service interactions