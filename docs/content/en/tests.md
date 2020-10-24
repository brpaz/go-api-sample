---
title: 'Tests'
position: 4
category: 'Application Structure'
fullscreen: false
---

## Unit tests

Unit tests helps guarantee that your code works as expected.

You can run then with `make test`.

This command will also generate a code coverage report that will be stored in `test/cover` directory.


## Integration Tests

Integration tests on this project are mostly used to test Database access (Repositories) with a real database.

These tests guarantees that our queries work as expected and that the migrations are well written.

You can execute this tests by running `make test-integration`.

The command will do:

* Connect to the existing database instance and create a test database.
* Run the application migrations.
* Runs the tests.
* Delete the test db.

This uses txdb, which encapsulates all the quries inside a transaction, making the tests isolated and without the need to any manual cleanup of test data between tests.

The test entrypoint is placed on `test/integration/db` directory.

## Acceptance Tests

Acceptance tests are higher level tests that guarantees the application features are working as expected.

We use godog to write the features in a BDD style and a separate test database.

You can see the features in `test/acceptance` directory.
