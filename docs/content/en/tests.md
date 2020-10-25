---
title: 'Testing'
position: 5
category: 'Code quality'
fullscreen: false
---

Testing and good coverage with all kinds of tests is essential to guarantee that the application works as expected.

## Unit tests

This project follows Go best practices of placing the tests in the same package as the code, with a `_test` suffix.

You can the application unit tests with `make test`. [Gotestsum](https://github.com/gotestyourself/gotestsum) is used for better output.

This command will also generate a code coverage report that will be stored in `test/cover` directory.


## Integration Tests

Integration tests are used to test code that interacts with external services like Databases.

In this project, we use integration tests to assert that "Repositories" queries work as expected against a real database instance, since doing unit tests for these wouldn't´t give us enough value.

We also test that our Migration scripts works.

You can execute this tests by running `make test-integration`.

The command will do:

* Connect to the existing database instance and create a test database.
* Run the application migrations.
* Runs the tests tagged with **integration**. [Learn more about build tags](https://mickey.dev/posts/go-build-tags-testing/)

The test entry point is located on `test/integration/db` directory.

<alert>
Make sure the application database container is running before executing the integration tests.
</alert>

The integration tests use the same `.env` file as the application to know the database connection settings.
By default, it will append `_test` to the value of `DB_DATABASE` environment variable and that will be the database that will be used to run the tests.


## Acceptance Tests

Acceptance tests are higher level tests that guarantees the application features are working as expected.

The tests are specified using the [Gherkin](https://cucumber.io/docs/gherkin/) language and executed with [godog](https://github.com/cucumber/godog).

You can see the features in `test/acceptance/features` directory.

To run these tests you can run `make test-acceptance`.

You can run this tests against an already running application by specifying `APP_URL` environment variable or let the command to start a new instance of the application and use the `httptest` server.

## Smoke Tests

Smoke tests are useful to check that an application is ready to serve requests and that the critical features are working. They are normally run after a deployment and against the real application instance.

This project provides a basic smoke test that you can run with `APP_URL=http://localhost:5000 make test-smoke`.

Like integration tests, it also uses Godog to run the tests.

## Container Tests

It´s not only the application code that needs to be tested. It´s also important to guarantee that any infrastructure related code is covered with tests.

Since we use Docker to deploy our application, the first step is to guarantee that our Dockerfile is correctly building the image.

For that we [Google Container Structure tests](https://github.com/GoogleContainerTools/container-structure-test) tool.

In these tests, we check that container image, contains the application binary and that is executable.

You can find the tests definitions in the `container-structure-test.yml` file at the root of the repository.

You can run these tests with: `make container-test`.

## Security Tests

Another important point of any application is security. 

We use [gosec](https://github.com/securego/gosec) to do Static Application Security Testing, integrated into [golang-ci](https://github.com/golangci/golangci-lint). See the [Linting](/linting) page to know more.

We should also include a Dynamic Application Security Testing tool like [OWASP Zap](https://owasp.org/www-project-zap/) at least in the CI Pipeline.

