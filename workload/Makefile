.phony: init-cluster start-cluster stop-cluster remove-cluster

init-cluster:
	./init-cluster.sh

start-cluster:
	k3d cluster start hlidskjalf

stop-cluster:
	k3d cluster stop hlidskjalf

remove-cluster:
	k3d cluster delete hlidskjalf