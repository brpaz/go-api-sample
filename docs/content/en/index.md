---
title: 'Introduction'
position: 1
category: 'Getting started'
fullscreen: false
---

This project provides a comprehensive example of how to build a complete Golang API, including some Clean architechture concepts, Dependency Injection, tests, CI Pipeline and more.

It demonstrates how I like to structure Go apps and aims to be my personal Golang API boilerplate.

## What is included

This project includes a simple API with endpoints to create and list todos. The idea is not to have a complex fully functional app but a boilerplate demonstrating architecture and practices.

It uses the following tools and libraries:

* [Echo](https://echo.labstack.com/) as the backbones of the application and HTTP Server.
* [Zap logger](https://github.com/uber-go/zap) for logging.
* [Gorm](https://gorm.io/index.html) for database access.
* [golang-migrate](https://github.com/golang-migrate/migrate) - Database migration tool.
* [sarulabs/di](https://github.com/sarulabs/di) as dependency injection framework.
* [godog](https://github.com/cucumber/godog) for writting acceptance tests using BDD.
* [GitHub Actions](https://github.com/features/actions) for the CI / CD Pipeline.
* [Docker](https://www.docker.com/) and [Docker-Compose](https://docs.docker.com/compose/) for easy development environment.
* [Nuxt Content](https://content.nuxtjs.org/) for this documentation site.

and more ...

## Running the project

To run this project, please see the [Installation Guide](/installation-guide).
