CLICKHOUSE_RELEASE=clickhouse

.phony: init-cluster start-cluster stop-cluster remove-cluster deploy-clickhouse deploy-pawn

init-cluster:
	./init-cluster.sh

start-cluster:
	k3d cluster start hlidskjalf

stop-cluster:
	k3d cluster stop hlidskjalf

remove-cluster:
	k3d cluster delete hlidskjalf

deploy-clickhouse:
	helm install $(CLICKHOUSE_RELEASE) clickhouse/chart

remove-clickhouse:
	helm install $(CLICKHOUSE_RELEASE)

deploy-pawn:
	kubectl apply -f $(pwd)/pawn/deployment.yaml
	kubectl apply -f $(pwd)/pawn/service.yaml