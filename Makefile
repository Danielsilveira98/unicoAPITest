setup:
	go mod download

test:
	go test -cover ./internal/...
