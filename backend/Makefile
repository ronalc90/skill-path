.PHONY: build run test clean db-up db-down docker-up docker-down lint

# Go parameters
BINARY_NAME=skillpath-api
MAIN_PATH=./cmd/server

## build: Build the Go binary
build:
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

## run: Run the application locally
run:
	go run $(MAIN_PATH)/main.go

## test: Run all tests
test:
	go test -v -race -cover ./...

## test-coverage: Run tests with coverage report
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

## clean: Remove build artifacts
clean:
	rm -rf bin/ coverage.out coverage.html

## db-up: Start PostgreSQL with Docker
db-up:
	docker compose up -d postgres

## db-down: Stop PostgreSQL
db-down:
	docker compose down

## docker-up: Start all services with Docker
docker-up:
	docker compose up --build -d

## docker-down: Stop all Docker services
docker-down:
	docker compose down -v

## lint: Run Go linter
lint:
	golangci-lint run ./...

## tidy: Tidy Go modules
tidy:
	go mod tidy

## help: Show available commands
help:
	@echo "Available commands:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | sort
