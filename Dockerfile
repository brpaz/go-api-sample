# Application dockerfile.
# - The build_base stage is responsible for preparing the envrionment and install go mod dependencies.
# - The dev stage is the stage using during development. it contains tools that are only needed in development like pg client, docker cli
# and custom entrypoints
# - The builder stage builds the application binary for production
# - The production stage only contains the application binary.
FROM golang:1.15-alpine AS build_base

ARG BUILD_DATE
ARG VCS_REF

SHELL ["/bin/ash", "-eo", "pipefail", "-c"]

# hadolint ignore=DL3019
RUN apk add bash curl ca-certificates git gcc g++ libc-dev make

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xz &&  \
	mv migrate.linux-amd64 /usr/local/bin/migrate && \
	migrate -version

RUN mkdir -p /src/app
WORKDIR /src/app
VOLUME /src/app

# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

#This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download
# command will _ only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod tidy

# Development stage
FROM build_base AS dev

SHELL ["/bin/ash", "-eo", "pipefail", "-c"]

# hadolint ignore=DL3019
RUN apk add postgresql-client

COPY --from=build_base /usr/local/bin/migrate /usr/local/bin/migrate

RUN curl -L https://github.com/vektra/mockery/releases/download/v2.3.0/mockery_2.3.0_Linux_x86_64.tar.gz  | tar xz -C /tmp &&  \
	mv /tmp/mockery /usr/local/bin/mockery && \
	mockery --version

RUN curl -L https://github.com/gotestyourself/gotestsum/releases/download/v0.5.3/gotestsum_0.5.3_linux_amd64.tar.gz | tar xz -C /tmp && \
	mv /tmp/gotestsum /usr/local/bin/gotestsum && \
	gotestsum --version

RUN go mod download

COPY docker/wait-for-it.sh /usr/local/bin/wait-for-it
COPY docker/docker-entrypoint.sh /usr/local/bin/docker-entrypoint

RUN chmod +x /usr/local/bin/wait-for-it && \
	chmod +x /usr/local/bin/docker-entrypoint

RUN go get github.com/markbates/refresh

ENTRYPOINT ["/usr/local/bin/docker-entrypoint"]

CMD ["refresh", "run"]

# This image builds the server for production usage
FROM build_base AS builder

# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/brpaz/go-api-sample/internal.BuildDate=${BUILD_DATE}' -X 'github.com/brpaz/go-api-sample/internal.BuildCommit=${VCS_REF}'"  -o /bin/server cmd/server/main.go

#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:3.11 AS production

ENV APP_ENV=prod

# Finally we copy the statically compiled Go binary.
COPY --from=builder /bin/server /bin/app
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

ENTRYPOINT ["/bin/app"]

LABEL maintainer="Bruno Paz <oss@brunopaz.dev>" \
      org.opencontainers.image.title="Go API Boilerplate" \
      org.opencontainers.image.description="Sample project demonstrating my best practices for a Golang based API" \
      org.opencontainers.image.url="https://github.com/brpaz/golang-api-sample" \
      org.opencontainers.image.source="git@github.com:brpaz/golang-api-sample" \
      org.opencontainers.image.vendor="Bruno Paz" \
      org.opencontainers.image.revision="$VCS_REF" \
      org.opencontainers.image.created="$BUILD_DATE"
