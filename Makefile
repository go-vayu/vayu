.PHONY: build run test lint lint-fix

build:
	@go build -o ./target/main ./main.go

run: build
	@./target/main

test:
	@go test ./...

lint:
	@golangci-lint run

lint-fix:
	@golangci-lint run --fix