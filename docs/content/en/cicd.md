---
title: 'CI/CD Pipeline'
position: 6
category: 'Application Structure'
fullscreen: false
---

[GitHub Actions](https://github.com/features/actions) is the tool we use for our CI Pipeline.

For each commit or Pull request, it will run a set of tasks to prepare our application to be deployed and to ensure it is complaint with our code quality metrics.

It contains the following jobs:

* **lint** - Runs linting tools like golangci-lint to ensure the quality of our code.
+ **unittests** - Runs unit tests
* **build** - Builds the application docker iamge.

...
