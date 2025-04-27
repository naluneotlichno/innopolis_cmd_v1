.PHONY: lint lint-setup migrate build build-docker run-local mockgen test

PROJECT_DIR=./chain-service

lint:
	@echo "Start linters..."
	@cd $(PROJECT_DIR) && golangci-lint run ./... && echo "All linters have run successfully"

lint-setup:
	chmod +x ./install-linter.sh && ./install-linter.sh

migrate:


build:
	go build -o App ./chain-service/cmd/chain-service/main.go

build-docker:
	docker-compose build

run-local:
	docker-compose up

mockgen:
	mockgen -source=chain-service/internal/repo/message_chain_repository.go \
	-destination=chain-service/internal/usecase/mocks/mocks_repo_test.go -package=usecase_test 
	mockgen -source=chain-service/internal/usecase/create_message_chain_usecase.go \
	-destination=chain-service/internal/usecase/mocks/mocks_usecase_test.go -package=usecase_test

test:
	cd $(PROJECT_DIR) && go test ./internal/usecase/mocks