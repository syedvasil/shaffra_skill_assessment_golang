MAKEFILE := $(abspath $(firstword $(MAKEFILE_LIST)))
MAKEFILE_DIR := $(abspath $(dir $(MAKEFILE)))
BIN_DIR := $(MAKEFILE_DIR)/bin

VERSION := $(shell git describe --always)
CURRENT_FILE := $(lastword $(MAKEFILE_LIST))

GOLANGCI_LINT_VERSION=v1.54.2

help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build binary
	@mkdir -p bin
	@go build  -o bin/${APP} ./cmd/server/

run: ## Run code
	@go run ./cmd/server/

generate: ## Generate mocks
	go generate ./...

install-tools: ## Install dev tools
	@go install github.com/vektra/mockery/v2@v2.20.0

godoc:
	@swag init -o ./docs -d ./cmd/server,./internal/app/controller/user --parseDependency

test: ## Run tests
	@go test ./... -race -coverprofile=./coverage/coverage.out -covermode=atomic

view_cov: ## show coverage
	@go tool cover -html=./coverage/coverage.out

coverage_ignore: ## Basic code coverage with text output
	@go test -covermode=count -coverprofile=./coverage/unfilteredcoverage.tmp $(shell go list ./... | egrep -v '(/test|/test/mock|/mocks)$$') && cat ./coverage/unfilteredcoverage.tmp | grep -v 'mock\|buggy' > ./coverage/coverage.out && go tool cover -func=./coverage/coverage.out

coverage-html: coverage_ignore
	@go tool cover -html=./coverage/coverage.out

## todo add linting
