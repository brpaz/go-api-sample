
DOCKER_LOCAL_IMAGE_NAME := "go-api:local-dev"
DOCKER_LOCAL_PROD_IMAGE_NAME := "go-api:local-prod"

# The run mode, helps running commands inside a Docker container or in local machine directly.
RUN_MODE_DOCKER:="docker"
RUN_MODE_LOCAL="local"
RUN_MODE ?= RUN_MODE_DOCKER

.PHONY: help fmt lint lint-dockerfile build build-docker test test-coverage test-acceptance start restart down logs sh
.DEFAULT_GOAL:=help
.SILENT: ;

-include .env

setup: ## Bootstraps the local development environment
	./scripts/setup-local-env.sh

fmt: ## Formats the go code using gofmt
	gofmt -w -s .

lint: ## Lint Go code
	docker run --rm -t -v $(PWD):/app -w /app golangci/golangci-lint:v1.30.0 golangci-lint run -v

lint-dockerfile: ## Lint Dockerfile
	docker run --rm -v "$(PWD):/data" -w "/data" hadolint/hadolint hadolint Dockerfile

gosec: ## Runs Gosec
	docker run --rm -t -v $(PWD):/app -w /app securego/gosec:v2.4.0 ./...

build: ## Build the app
	go build -o build/server cmd/server/main.go

build-docker: ## Builds the application using Docker for development purposes
	docker build  \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
        --build-arg VCS_REF="" \
		--target=dev \
		. -t $(DOCKER_LOCAL_IMAGE_NAME)

build-docker-prod: ## Builds the application using Docker for production purposes
	docker build  \
    		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
            --build-arg VCS_REF="" \
    		--target=production \
    		. -t $(DOCKER_LOCAL_PROD_IMAGE_NAME)

test: ## Run unit tests
ifeq ($(RUN_MODE), RUN_MODE_DOCKER)
	make build-docker
	docker run --rm -t -v "$(PWD):/src/app" $(DOCKER_LOCAL_IMAGE_NAME) go test -v -short ./...
else
	go test -v -race -short ./...
endif

test-coverage: ## Run tests with coverage
	mkdir -p test/cover
	go test -v -race -short -coverprofile ./test/cover/cover.out -covermode=atomic  ./...
	go tool cover -html=./test/cover/cover.out

test-acceptance: ## Runs acceptance tests
	go test -v ./test/functional/functional_test.go -godog.random -tags=functional -count=1

start: ## Starts the application
	docker-compose up -d

restart: ## Restarts the application
	docker-compose restart app

down: ## Stops the applications
	docker-compose down

logs: start ## Display logs of the application
	docker-compose logs -f app

sh: start ## Opens a terminal in the application container
	docker-compose exec app sh

container-test: build-docker-prod ## Runs container structure tests
	docker run -i --rm \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -v $(PWD):/test -w /test zemanlx/container-structure-test:v1.8.0-alpine \
        test \
        --image $(DOCKER_LOCAL_PROD_IMAGE_NAME) \
        --config container-structure-test.yaml

help: ## Displays help menu
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
