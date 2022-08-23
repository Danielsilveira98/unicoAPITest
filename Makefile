#!make
include .env
export $(shell sed 's/=.*//' .env)

setup:
	go mod download

test:
	go test -cover ./internal/...

lint:
	golangci-lint run ./...

updb:
	POSTGRES_USER=${POSTGRES_USER} POSTGRES_PASSWORD=${POSTGRES_PASSWORD} docker-compose up -d postgres && docker-compose logs -f postgres

stopdb:
	docker-compose stop postgres

loadfiles:
	go run ./scripts/populate_db/script.go > log/populate_db.log

run:
	go run cmd/api/main.go


getall:
	curl http://localhost:8000/street_market
