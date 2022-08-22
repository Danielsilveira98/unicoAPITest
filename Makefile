setup:
	go mod download

test:
	go test -cover ./internal/...

lint:
	golangci-lint run ./...

updb:
	docker-compose up -d postgres && docker-compose logs -f postgres

stopdb:
	docker-compose stop postgres

loadfiles:
	MIGRATIONS_PATH=deployment/migrations DATA_PATH=scripts/populate_db/data/ go run ./scripts/populate_db/script.go > log/populate_db.log

run:
	go run cmd/main.go
