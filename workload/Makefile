.phony: init-cluster start-cluster stop-cluster remove-cluster

init-cluster:
	./scripts/init-cluster.sh

start-cluster:
	k3d cluster start hlidskjalf

stop-cluster:
	k3d cluster stop hlidskjalf

remove-cluster:
	k3d cluster delete hlidskjalf

set-context:
	kubectl config set-context --current --namespace=elysium

start-deployments:
	./scripts/start-deployments.sh

start-servies:
	./scripts/start-services.sh

stop-deployments:
	./scripts/stop-deployments.sh

stop-services:
	./scripts/stop-services.sh	
