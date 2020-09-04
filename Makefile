# =======================================================================================
# This makefile contains common tasks to help with the development of the application.
# =======================================================================================

# Make configurations
.PHONY: help fmt lint lint-dockerfile build build-docker test test-coverage test-acceptance start restart down logs sh
.DEFAULT_GOAL:=help
.SILENT: ;

# Project variables
PROJECT_NAME:="go-api"

DOCKER_LOCAL_IMAGE_NAME:= $(PROJECT_NAME):local-dev
DOCKER_CONTAINER_TESTS_IMAGE:=$(PROJECT_NAME):container-tests

NOW:=`date -u +"%Y-%m-%dT%H:%M:%SZ"`
VERSION:=$(shell git rev-parse --short HEAD)-dev

# The run mode, helps running commands inside a Docker container or in local machine directly.
RUN_MODE_DOCKER:="docker"
RUN_MODE_LOCAL="local"
RUN_MODE ?= RUN_MODE_LOCAL

-include .env

setup: ## Bootstraps the local development environment
	./scripts/setup-local-env.sh

fmt: ## Formats the go code using gofmt
	gofmt -w -s .

lint: ## Lint Go code
	docker run --rm -t -v $(PWD):/app -w /app golangci/golangci-lint:v1.30.0 golangci-lint run -v

lint-dockerfile: ## Lint Dockerfile
	docker run --rm -v "$(PWD):/data" -w "/data" hadolint/hadolint hadolint Dockerfile

build: ## Build the app
	go build -o build/app -ldflags="-X 'github.com/brpaz/go-api-sample/internal/buildinfo.BuildDate=$(NOW)' -X 'github.com/brpaz/go-api-sample/internal/buildinfo.BuildCommit=$(VERSION)'" cmd/server/main.go

run: build ## Runs the app
	./build/app

build-docker: ## Builds the application using Docker for development purposes
	docker build  \
		--build-arg BUILD_DATE=$(NOW) \
        --build-arg VCS_REF=$(VERSION) \
		--target=dev \
		. -t $(DOCKER_LOCAL_IMAGE_NAME)

test: ## Run unit tests
ifeq ($(RUN_MODE), RUN_MODE_DOCKER)
	make build-docker
	docker run --rm -t -v "$(PWD):/src/app" $(DOCKER_LOCAL_IMAGE_NAME) go test -v -short ./...
else
	go test -v -race -short  -coverprofile ./test/cover/cover.out -covermode=atomic  ./...
	go tool cover -html=./test/cover/cover.out
endif

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

container-test: ## Runs container structure tests on Docker image
	docker build  \
    		--build-arg BUILD_DATE=$(NOW) \
            --build-arg VCS_REF=$(VERSION) \
    		--target=production \
    		. -t $(DOCKER_CONTAINER_TESTS_IMAGE)

	docker run -i --rm \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -v $(PWD):/test -w /test zemanlx/container-structure-test:v1.8.0-alpine \
        test \
        --image $(DOCKER_CONTAINER_TESTS_IMAGE) \
        --config container-structure-test.yaml

help: ## Displays help menu
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
