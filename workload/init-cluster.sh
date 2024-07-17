# remove exists cluster
k3d cluster delete hlidskjalf
# create new cluster
k3d cluster create hlidskjalf
# add more agents
k3d node create worker-1 --cluster=hlidskjalf --role=agent
k3d node create worker-2 --cluster=hlidskjalf --role=agent
k3d node create worker-3 --cluster=hlidskjalf --role=agent