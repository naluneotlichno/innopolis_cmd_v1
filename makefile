.PHONY: lint lint-setup migrate build build-docker run-local test test-mockss

# Проверка линтером
lint:
	@echo "Start linters..."
	@golangci-lint run ./... && echo "All linters have run successfully"

lint-setup:
	chmod +x ../install-linter.sh && ../install-linter.sh

migrate:


build:
	go build -o App ./chain-service/cmd/chain-service/main.go

build-docker:
	docker-compose build

run-local:
	docker-compose up

test:
	go test ./...

test-mocks:

