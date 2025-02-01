.PHONY: gen
gen:
	buf lint
	buf format
	buf generate

.PHONY: up
up:
	docker-compose -f ./docker/docker-compose.local.yaml up -d

.PHONY: down
down:
	docker-compose -f ./docker/docker-compose.local.yaml down
