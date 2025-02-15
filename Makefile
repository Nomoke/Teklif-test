.PHONY: run test build docker lint migrate

BINARY=teklifhub

GO_ENV=development
DATABASE_URL=postgres://user:password@localhost:5432/teklifhub

run:
	go run cmd/app/main.go

build:
	go build -o $(BINARY) cmd/app/main.go

migrate:
	docker-compose -f deploy/docker-compose.yaml --env-file deploy/.env up migrate

docker-up:
	docker-compose -f deploy/docker-compose.yaml --env-file deploy/.env up --build

docker-down:
	docker-compose -f deploy/docker-compose.yaml down
