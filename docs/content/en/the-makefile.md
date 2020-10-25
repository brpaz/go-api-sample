---
title: 'The Makefile'
position: 4
category: 'Getting started'
fullscreen: false
---

The application comes with a `Makefile` that provides tasks for common needs like starting the application, inspecting the logs, running migrations etc.

You can see the available tasks by running `make help` on the terminal. some include:

* **fmt** - Runs `go fmt` to format your code.
* **lint** - Runs `golangci-lint` to lint your code.
* **test**  - Runs unit tests.
* **test-integration** - Runs integration tests
* **test-acceptance** - Runs acceptance tests.
* **build** - Builds the application using Docker Compose.
* **up** - Starts application containers using Docker Compose.
* **down** - Stops the application containers
* **sh** - Opens a shell to the application container.
* **logs**  - Show container logs.

