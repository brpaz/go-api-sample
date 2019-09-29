# Base build image
FROM golang:1.13-alpine AS build_base

# Install some dependencies needed to build the project

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
RUN go get github.com/markbates/refresh

CMD refresh run

# This image builds the server for production usage
FROM build_base AS builder

# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/server cmd/server/main.go

#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:3.10 AS final

# Finally we copy the statically compiled Go binary.
COPY --from=builder /bin/server /bin/app

ENTRYPOINT ["/bin", "app"]
