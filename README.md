# go-api-sample

> Demo project showing a full CI /CD Pipeline for a Golang Application using GitHub Actions and Google Cloud

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)

[![GitHub Actions](https://github.com/brpaz/go-api-sample/workflows/CI/badge.svg?style=for-the-badge)](https://github.com/brpaz/go-api-sample/actions)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/8c79d720eb364a2cb0ef2f3d98a1874d)](https://www.codacy.com/manual/brpaz/go-api-sample?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=brpaz/go-api-sample&amp;utm_campaign=Badge_Grade)

## What is Included

* API based on [Echo framework](https://echo.labstack.com/)
* Logging with [Zap](https://github.com/uber-go/zap)
* Hot Reload with [Refresh](https://github.com/markbates/refresh)
* DotEnv support
* Docker and Docker-compose
* Full CI/CD pipeline using GitHub Actions
* Deploy with [Google Cloud run](https://cloud.google.com/run/)
* Functional Tests with [godog](https://github.com/DATA-DOG/godog)
* Code quality metrics with [Codacy](https://codacy.com)

## Pre-requisites

- Go 1.14 or higher with [go modules support](https://github.com/golang/go/wiki/Modules) enabled.

## Usage

The recommended way is to use [Docker](https://www.docker.com/).

You can start the application with ```docker-compose up```command. A Makefile is provided in the root of the repository with useful tasks like running tests.

## Todo

* Add semantic release support.


## 🤝 Contributing

Contributions, issues and feature requests are welcome!

## Author

👤 **Bruno Paz**

  * Website: [https://github.com/brpaz](https://github.com/brpaz)
  * Github: [@brpaz](https://github.com/brpaz)

## 📝 License

Copyright © 2020 [Bruno Paz](https://github.com/brpaz).

This project is [MIT](LICENSE) licensed.
