.PHONY: build run test clean migrate-up migrate-down help

BINARY_NAME=bookapi
DB_DSN="postgres://postgres:2115@localhost:5432/booksdb?sslmode=disable"

build:
	go build -o $(BINARY_NAME) ./cmd/api/main.go

run: build
	./$(BINARY_NAME)

test:
	go test -v ./...

clean:
	rm -f $(BINARY_NAME)

migrate-up:
	migrate -path ./migrations -database $(DB_DSN) up

migrate-down:
	migrate -path ./migrations -database $(DB_DSN) down

help:
	@echo "Available commands:"
	@echo "  make build        - build binary"
	@echo "  make run          - build and run"
	@echo "  make test         - run tests"
	@echo "  make clean        - remove binary"
	@echo "  make migrate-up   - apply database migrations"
	@echo "  make migrate-down - rollback migrations"