---
title: 'Packages Overview'
position: 2
category: 'Application Structure'
fullscreen: false
---

As described in the [Directory Structure](/directory-structure) page, the application code is stored in the `internal` folder and.

This package is then organized in sub-packages according to the application responsibilities:

```
internal
├── app
├── config
├── db
├── errors
├── http
├── logging
├── todo
├── util
└── validator
```

## app

The `app` package contains the code needed to bootstrap the application, like DI Configuration and HTTP Server bootstrap.

Check [The application](/structure/app) for more details.

## config

This package contains the definition of the struct that will hold the application configuration as well as methods to create a new instance of that struct with values from environment variables.

You can read more about the configuration loading process [here](/structure/config)

## db

This package ackage contains code to initialize and manage Database connections.

## errors

This package contains the definition of error codes and custom error types used across the application.

## logging

This package contains code that is responsible to initialize the application logger. Look [here](/structure/logging) for details.

## http

Ths package provides common code related to http requests like middelwares, error handlers and global handlers like healhcheck.

The routes definitions could be placed in this package as well but since it´s the entrypoint to the web server and it needs to access DI I placed in `app` package, but this might change in the future.

Ex: Change the handler related service names from DI to http package and pass a list of built handlers as a map as well as the `echo` instance to the registerRoutes funcion.


## util

Miscellaneous utility functions used across the application.

## validator

Custom validators to be used across the application.

## todo

This package contains the code that implements the todo functionality. 

Each module of the application will have it´s own package containing everything needed to impelement the feature including http handlers, application services and repositories.

This code follows some patterns from [Clean Architecture](https://www.freecodecamp.org/news/a-quick-introduction-to-clean-architecture-990c014448d2/) and Hexagonal Architecture.

For this project we place all the files together in the package, but you can further create subpackages like `handlers`, `usecases`, `domain`, etc for more cleaner separation of concerns.
