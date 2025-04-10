.PHONY: lint lint-setup migrate build build-docker run-local test test-mocks

lint:
	golangci-lint run ./...

lint-setup:
	chmod +x install-linter.sh && ./install-linter.sh

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

