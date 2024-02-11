.PHONY = start-registry-infra stop-registry-infra

start-registry-infra:
	docker-compose -f ./harbor-registry/docker-compose.registry.infra.yaml --env-file .env.docker up -d

stop-registry-infra:
	docker-compose -f ./harbor-registry/docker-compose.registry.infra.yaml --env-file .env.docker down