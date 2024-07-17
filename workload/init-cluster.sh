# remove exists cluster
k3d cluster delete hlidskjalf
# create new cluster
k3d cluster create hlidskjalf
# add more agents
k3d node create worker-1 --cluster=hlidskjalf --role=agent
k3d node create worker-2 --cluster=hlidskjalf --role=agent
k3d node create worker-3 --cluster=hlidskjalf --role=agent
# taint master
kubectl taint nodes k3d-hlidskjalf-server-0 node-role.kubernetes.io/master=:NoSchedule
# label workers
kubectl label node k3d-worker-1-0 node-role.kubernetes.io/worker=worker
kubectl label node k3d-worker-2-0 node-role.kubernetes.io/worker=worker
kubectl label node k3d-worker-3-0 node-role.kubernetes.io/worker=worker