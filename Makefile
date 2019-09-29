
.PHONY: help env dep fmt lint vet build build-docker test cover version
.DEFAULT: help

setup: ## Boostraps the application
	[ ! -f .env ] && cp .env.example .env || true
	GO111MODULE=off go get github.com/markbates/refresh github.com/mgechev/revive
	pre-commit install

dep: ## Get build dependencies
	go mod tidy

fmt: ## Formats the go code using gofmt
	gofmt -w -s .

lint: ## Lint code
	revive -config revive.toml -formatter friendly ./...

vet: ## Run go vet
	go vet ./...

build: ## Build the app
	@go build -o build/server cmd/server/main.go

test: ## Run package unit tests
	@go test -v -race -short ./...

test-coverage: ## Run tests with coverage
	mkdir -p test/cover
	@go test -v -race -short -coverprofile ./test/cover/cover.out -covermode=atomic  ./...

test-functional: ## Runs functional
	go test -v ./test/functional/functional_test.go -godog.random -tags=functional -count=1

up: ## Starts the application
	docker-compose up

server: ## Starts the application server
	refresh run

help: ## Displays help menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
