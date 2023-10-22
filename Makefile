# Makefile for the pingpong project

.PHONY: build test lint clean run coverage

## build: Compiles the main application.
build:
	@echo "Building pingpong..."
	@go build -o pingpong ./cmd/pingpong

## test: Runs the unit tests.
test:
	@echo "Running tests..."
	@go test ./... -cover

## clean: Removes binary and coverage files.
clean:
	@echo "Cleaning up..."
	@rm -f pingpong
	@rm -f coverage.out

## run: Runs the application.
run: build
	@echo "Starting pingpong..."
	@./pingpong

## coverage: Generates test coverage report.
coverage:
	@echo "Generating coverage report..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

## help: Show this help message.
help:
	@echo "Choose a command run in pingpong:"
	@echo ""
	@echo "Usage:"
	@echo "  make <command>"
	@echo ""
	@echo "Available commands:"
	@sed -n 's/^##//p' Makefile | column -t -s ':' |  sed -e 's/^/  /'
