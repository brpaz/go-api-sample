---
title: 'Directory Structure'
position: 3
category: 'Getting started'
fullscreen: false
---

This project directory structure, follows some practices of the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) like using **internal** to organize application related code or **cmd** package to place the entry points of the application.

Here is how the top-level project directory looks like:

```
.
├── cmd
├── docker
├── docs
├── internal
├── migrations
├── scripts
└── test
```

### The cmd directory

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, put that code in the /internal directory. You'll be surprised what others will do, so be explicit about your intentions!

### The Internal directory

In this directory, goes the private application and library code, that you don't want others importing in their applications or libraries. 

Since this is an application and not an importable library, makes sense to have all the code in the internal folder.

The application code inside the internal directory is further organized into modular packages. You can see how it´s organized in the [Application Packages](/app-pakcages) page.


### Docker

This folder contains files that are used in the Docker image, like entry point scripts or configuration scripts (Ex: nginx.conf).

The Dockerfile and Docker-Compose file are still placed in the root of the repository.

### Docs

This directory Contains the sources for the project documentation. The Open API definition for the API can be found in the "api" folder and this documentation sources in Markdown files in the "content" folder.

We use [Nuxt Content](https://content.nuxtjs.org/) and [Redoc](https://github.com/Redocly/redoc) to generate documentation from this, so the docs folder also have files needed for this tools to work, like `nuxt.config.js` and `package.json`.

### Migrations

Contains the database migrations sql scripts. We use [golang-migrate](https://github.com/golang-migrate/migrate) to manage the application migrations.

### Scripts

Utility bash scripts that can be used in CI or in the local environment.

### Test

This directory contains integration and acceptance tests code as well as global utility functions for testing used across the project.

The unit tests are placed in the same package to the files under test, following Go recommendations.

You can read more about our testing strategy in the [Testing](/tests) page of this documentation.

### Other files.

We have a couple of dotfiles and configuration files in the root of the application like:

* .env - Defines envrionment variables for the application
* .czrc - [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) configuration
* .golangci.yml - Configuration for the linter tool [golangci](https://github.com/golangci/golangci-lint)
* Dockerfile - Project Dockerfile
* Makefile - Common tasks like starting the application or running tests.

And more.
