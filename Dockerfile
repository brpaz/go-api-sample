FROM golang:1.14-alpine AS build_base

ARG BUILD_DATE
ARG VCS_REF

# hadolint ignore=DL3019
RUN apk add bash ca-certificates git gcc g++ libc-dev make

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
RUN go mod download

RUN go get github.com/markbates/refresh

CMD ["refresh", "run"]

# This image builds the server for production usage
FROM build_base AS builder

# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/brpaz/go-api-sample/internal/buildinfo.BuildDate=${BUILD_DATE}' -X 'github.com/brpaz/go-api-sample/internal/buildinfo.BuildCommit=${VCS_REF}'"  -o /bin/server cmd/server/main.go

#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:3.11 AS production

ENV APP_ENV=prod

# Finally we copy the statically compiled Go binary.
COPY --from=builder /bin/server /bin/app

ENTRYPOINT ["/bin/app"]

LABEL maintainer="Bruno Paz <oss@brunopaz.dev>" \
      org.opencontainers.image.title="Go API Sample" \
      org.opencontainers.image.description="Sample project demonstrating my best practices for a Golang based API" \
      org.opencontainers.image.url="https://github.com/brpaz/golang-api-sample" \
      org.opencontainers.image.source="git@github.com:brpaz/golang-api-sample" \
      org.opencontainers.image.vendor="Bruno Paz" \
      org.opencontainers.image.revision="$VCS_REF" \
      org.opencontainers.image.created="$BUILD_DATE"
