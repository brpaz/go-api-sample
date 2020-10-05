---
title: 'The Makefile'
position: 3
category: 'Application Structure'
fullscreen: false
---

The application comes with a Makefile that provides tasks for common needs like starting the application, inspecting the logs, running migrations etc.

You can inspect the Makefile or run `make help` to see the available tasks. some include:

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

And more ...
