.PHONY: lint setup

lint:
	golangci-lint run ./...

setup:
	chmod +x install-linter.sh && ./install-linter.sh