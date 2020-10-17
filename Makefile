# ====================================================================
# This makefile contains common tasks to help with the development
# of the application.
# ====================================================================

.PHONY: help fmt lint lint-docker test test container-test build up down restart db-migrate db-migration-create logs sh docs docs-api mocks
.DEFAULT_GOAL:=help
.SILENT: ;

PROJECT_NAME:="go-api"
DOCKER_CONTAINER_TESTS_IMAGE:=$(PROJECT_NAME):container-tests
APP_CONTAINER_NAME=app

NOW:=`date -u +"%Y-%m-%dT%H:%M:%SZ"`
VERSION:=$(shell git rev-parse --short HEAD)-dev
CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

targets := $(lastword $(MAKEFILE_LIST))

-include .env # loads the .env file

COMPOSE_RUN := docker-compose run --entrypoint "" $(APP_CONTAINER_NAME)
setup: ## Bootstraps the local development environment
	./scripts/setup-local-env.sh

# ==============================================================
# Test and Lint commands
# ==============================================================

fmt: ## Formats the go code using gofmt
	docker run --rm -t -v $(PWD):/app -w /app golang:1.14-alpine gofmt -w -s .

lint: lint-go lint-docker ## Runs all the linters

lint-go: ## Lint Go code
	docker run --rm -t -v $(PWD):/app -w /app golangci/golangci-lint:v1.30.0 golangci-lint run -v

lint-docker: ## Lint Dockerfile
	docker run --rm -v "$(PWD):/data" -w "/data" hadolint/hadolint hadolint Dockerfile

test: ## Run unit tests
	$(COMPOSE_RUN) gotestsum --format testname -- -v -tags=unit -coverprofile ./test/cover/cover.out -covermode=atomic  ./...

test-integration: ## Runs acceptance tests
	docker-compose run --entrypoint "" -e APP_ENV=test $(APP_CONTAINER_NAME) go run test/integration/db/main.go

test-acceptance: ## Runs acceptance tests
	$(COMPOSE_RUN) go test -v ./test/acceptance -tags=acceptance -count=1  -godog.format=pretty -godog.tags="$(TAGS)"

test-smoke: ## Runs smoke tests
	docker-compose exec -e APP_ENV="test" -e APP_URL=http://localhost:5000 $(APP_CONTAINER_NAME) go test -v ./test/smoke -godog.random -godog.format=pretty -tags=smoketests -count=1

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

# ==============================================================
# Docker-Compose commands
# ==============================================================
build: ## Builds the application using Docker Compose
	docker-compose build $(APP_CONTAINER_NAME)

up: ## Starts the application containers
	docker-compose up -d

restart: ## Restarts the application
	docker-compose restart $(APP_CONTAINER_NAME)

down: ## Stops the applications
	docker-compose down

logs: ## Display logs of the application
	docker-compose logs -f

sh: ## Opens a terminal in the application container
	docker-compose exec $(APP_CONTAINER_NAME) bash

db-migration-create: ## Creates a new Migration file
	$(COMPOSE_RUN) migrate create -ext sql -dir migrations $NAME
	sudo chown -R $(CURRENT_UID):$(CURRENT_GID) migrations

db-migrate: ## Runs Database migrations
	$(COMPOSE_RUN) migrate -path=migrations -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable up

mocks: ## Generate Mocks
	$(COMPOSE_RUN) mockery --all --inpackage --case underscore

# ==============================================================
# Other
# ==============================================================

docs: ## Render Application Documentation
	cd docs && yarn dev

docs-api: ## Render API documentation
	cd docs && yarn redoc:serve

help: ## Displays help menu
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' "Makefile" | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
