---
title: 'Installation guide'
position: 2
category: 'Getting started'
fullscreen: false
---

## Pre-Requisites

Before starting, make sure you have the following tools installed on your machine:

* Git
* Docker
* Docker-Compose
* Make
* Python (For pre-commit hooks)

This project ships a Docker-compose file and a Makefile with common tasks, to facilitate the development process.

## Setup

The first step is clone the project using git to a directory of your choice.

```shell script
git clone https://github.com/brpaz/go-api-sample
cd go-api-sample
```

After that, run:

```shell script
make setup
```

This will run some setup tasks like installing [Pre-commit hooks](https://pre-commit.com).

## Running the application

To start the application, just run:

```shell script
make up
```

This make task will run docker-compose under the hood to build and launch the application containers. It will launch 3 containers:

* App - With the application code
* Postgres - for the database
* Nginx-Proxy - for gateway

You can run `m̀ake logs` to see the application logs.

## Access the application.

There are two ways you can access the application, directly by the container IP address or by domain name.

In a more complex project with lots of services, it´s recommended using the domain name approach.

### Access by domain name.

By default, the application will be listening on `go-api.docker` domain.
You must add an entry to your `/etc/hosts` pointing this domain to `127.0.0.1`  or configure a DNS server to
point the domain to `127.0.0.1`.

After that, you can access the application with the following url:

```http://go-api.docker:8080/_health```

<alert>
8080 is the port that Nginx Proxy is configured to listen to. The internal application port is not exposed.

You can change this port, by specifying **NGINX_PROXY_PORT** in the .env file.

</alert>

### Access by container IP

You can find the container IP, by running `docker-compose ps` and inspecting the "Ports" section of the output.

For example, `0.0.0.0:32768->5000/tcp` indicates the application is listening on port 32768 as the internal port is never exposed.

You can then access the application, using `http://localhost:32768`.
