---
title: 'Code Quality'
position: 4
category: 'Application Structure'
fullscreen: false
---

## Linting / Static Analysis

There are many tools in the Go ecosystem that can be used to ensure our code meets quality standards, like go vet,

[Golangci-lint](https://github.com/golangci/golangci-lint) encapsulates many of these tools in a single tool and so I have included it in this project.

You can run it using `make lint`.

Some other linters are also available, like [Hadolint](https://github.com/hadolint/hadolint), which lints the application dockerfile.
