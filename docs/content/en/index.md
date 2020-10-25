---
title: 'Introduction'
position: 1
category: 'Getting started'
fullscreen: false
---

This project provides an example and boilerplate of a Golang API, demonstrating  some best practices like Clean architechture concepts, Dependency Injection, tests, CI/CD and more.

## What is included

This project includes a simple API with endpoints to create and list todos. The idea is not to have a complex fully functional app but a boilerplate demonstrating architecture and practices.

You can find the API documentation [here](/apidoc)

It uses some popular libraries and tools like:

* [Echo](https://echo.labstack.com/) - HTTP Server
* [Zap logger](https://github.com/uber-go/zap) - Logging.
* [Gorm](https://gorm.io/index.html) - Database access (ORM)
* [golang-migrate](https://github.com/golang-migrate/migrate) - Database migration tool.
* [sarulabs/di](https://github.com/sarulabs/di) - dependency injection framework.
* [godog](https://github.com/cucumber/godog) - BDD Framework for writing integration tests
* [GitHub Actions](https://github.com/features/actions) - CI/CD Pipeline.
* [Docker](https://www.docker.com/) and [Docker-Compose](https://docs.docker.com/compose/) - development environment.
* [Nuxt Content](https://content.nuxtjs.org/) for this documentation site.
* [Recdoc](https://github.com/Redocly/redoc) - API documentation 

## Running the project

To run this project, please see the [Installation Guide](/installation-guide).
