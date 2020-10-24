---
title: 'Directory Structure'
position: 2
category: 'Application Structure'
fullscreen: false
---

This application follows some best practices like using "internal" and "cmd" packages to structure go code and some concepts of [Hexagonal/Clean Architecture](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/) like use cases to write more modular code.

## Application Code

* **cmd** - This folder contains all the entrypoints for the application. In this example, only the web server (cmd/server) exists, but if you want to build a cli or some other utility commands, you can place them here.2
* **internal** This is where the application go code is placed.
* **test** Go best practices defines the test files together with the source files they are testing and we use that for unit tests. But since integration and acceptance tests require some bootstrap first (using TestMain), we use this folder to place the code for that.

## Other

* **docker** - Contains Docker specific files and scripts like the application entrypoint.
* **docs** - Project documentation goes in this folder.
* **scripts** - This is the place for any scripts, like to be running on CI or local.
