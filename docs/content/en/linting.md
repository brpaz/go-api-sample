---
title: 'Linting And Static Analysis'
position: 4
category: 'Code quality'
fullscreen: false
---

## Linting Go code

There are many tools in the Go ecosystem that can be used to ensure the code meets quality standards, like [go lint](https://github.com/golang/lint) [go vet](https://golang.org/cmd/vet/), [go cyclo](https://github.com/fzipp/gocyclo) and many more.

Instead of running each tool in separate we use [Golangci-lint](https://github.com/golangci/golangci-lint) which encapsulates this and many more linters and run them with a single command and in parallel.

You can run golangci-lint with `make lint-go` command.

For code format, we use the standard [gofmt](https://golang.org/cmd/gofmt/) that you can run with `make fmt`.

## Linting Docker

Not only the application code benefits from linting. [Hadolint](https://github.com/hadolint/hadolint) allows to verify the quality of our Dockerfile.

You can run hadolint with `make lint-docker`.
