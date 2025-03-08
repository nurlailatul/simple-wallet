.PHONY: run build clean test fmt vet

# Variables
APP_NAME=domper #dompet-paper

run: build
	@echo "Starting the app..."
	@./$(APP_NAME) start

migrate: build
	@echo "Migrate up..."
	@./$(APP_NAME) migrate:run

migrate-down: build
	@echo "Migrate down..."
	@./$(APP_NAME) migrate:rollback

build:
	@echo "Building the binary..."
	@go build -o $(APP_NAME)

test:
	@echo "Running tests..."
	@go test ./... -v

fmt:
	@echo "Formatting code..."
	@go fmt ./...

vet:
	@echo "Running go vet..."
	@go vet ./...

deps:
	@echo "Downloading dependencies..."
	@go mod tidy
