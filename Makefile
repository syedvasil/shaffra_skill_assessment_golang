GOLANGCI_LINT_VERSION=v1.54.2

build: ## Build binary
	@mkdir -p bin
	@go build  -o bin/${APP} ./cmd/server/

run: ## Run code
	@go run ./cmd/server/

install-tools: ## Install dev tools
	@go install github.com/vektra/mockery/v2@v2.20.0

godoc:
	@swag init -o ./docs -d ./cmd/server,./internal/app/controller/post,./internal/app/controller/user --parseDependency

test: ## Run tests
	@go test ./... -race -coverprofile=./coverage/coverage.out -covermode=atomic

coverage: ## show coverage
	@go tool cover -html=./coverage/coverage.out


## todo add linting
