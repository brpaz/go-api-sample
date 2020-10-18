# go-api-sample

> Comprehensive example of how to build a complete Golang API, including Clean architecture concepts, Dependency Injection, tests, CI Pipeline, documentation and more.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)
[![GitHub Actions](https://github.com/brpaz/go-api-sample/workflows/CI/badge.svg?style=for-the-badge)](https://github.com/brpaz/go-api-sample/actions)


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

## Getting started

To understand how this application is structured and how you can run it on your local machine, please see the [documentation](https://brpaz.github.io/go-api-sample/).


## Author

👤 **Bruno Paz**

  * Website: [https://github.com/brpaz](https://github.com/brpaz)
  * Github: [@brpaz](https://github.com/brpaz)

## 📝 License

Copyright © 2020 [Bruno Paz](https://github.com/brpaz).

This project is [MIT](LICENSE) licensed.
