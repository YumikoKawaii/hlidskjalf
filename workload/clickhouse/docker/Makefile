.phony: start stop

start:
	docker-compose --env-file .env.docker -f docker-compose.yaml up -d

stop:
	docker-compose --env-file .env.docker -f docker-compose.yaml down