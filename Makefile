.PHONY: build run test lint lint-fix generate-swagger-docs

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

generate-swagger-docs:
	@swag init -g ./internal/api/routes/routes.go -o ./internal/api/swagger