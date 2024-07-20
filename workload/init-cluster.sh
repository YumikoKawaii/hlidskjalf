## remove exists cluster
k3d cluster delete hlidskjalf

## create new cluster
k3d cluster create hlidskjalf

## add more agents
k3d node create worker-1 --cluster=hlidskjalf --role=agent
k3d node create worker-2 --cluster=hlidskjalf --role=agent
k3d node create worker-3 --cluster=hlidskjalf --role=agent

## limiting resources for nodes
# docker update --memory=1024m --memory-swap=2g --cpu-shares=128 k3d-hlidskjalf-server-0
# docker update --memory=128m --memory-swap=256m --cpu-shares=128 k3d-hlidskjalf-serverlb
# docker update --memory=2048m --memory-swap=4g --cpu-shares=512 k3d-worker-1-0
# docker update --memory=2048m --memory-swap=4g --cpu-shares=512 k3d-worker-2-0
# docker update --memory=2048m --memory-swap=4g --cpu-shares=512 k3d-worker-3-0

## taint master
kubectl taint nodes k3d-hlidskjalf-server-0 node-role.kubernetes.io/master=:NoSchedule

## label workers
kubectl label node k3d-worker-1-0 node-role.kubernetes.io/worker=worker
kubectl label node k3d-worker-2-0 node-role.kubernetes.io/worker=worker
kubectl label node k3d-worker-3-0 node-role.kubernetes.io/worker=worker

kubectl create namespace pawn