---
title: 'CI/CD Pipeline'
position: 7
category: 'Other'
fullscreen: false
---

Since the code is hosted at [Github](https://github.com), it makes a lot of sense to use [GitHub Actions](https://github.com/features/actions) for our CI Pipeline.

For each commit or Pull request, it will run a set of tasks to prepare our application to be deployed and to ensure it is compliant with our code quality metrics.

Our pipeline consists of three workflows:

* CI - Which is responsible to build and test the application code and push the built Docker image to a container registry.
* Build Docs - Which builds the application documentation and deploy to GitHub Pages.
* Release - Which can be triggered to release a new version of the application.


### CI

CI is the main Pipeline to build the application. It includes the following jobs:

#### lint 

Runs [golang-ci](https://github.com/golangci/golangci-lint) linter tool to lint go code.

### lint-docker 

Runs [Hadolint](https://github.com/hadolint/hadolint) to Lint the Dockerfile.

### test-unit

Runs unit tests and uploads Code coverage reports to [Codecov](https://codecov.io/) and as Build Artifacts.

### test-integration

Runs integration tests against a Postgres service.

### build-image

After we have verified that the code is working, this stage builds a Docker image from it.

### test-acceptance.

Fetch the Docker image created in the previous stage and runs acceptance tests against it, by starting the application.

### push-image

After all the tests are complete, it pushed the image to the Docker registry, in this project, [Docker hub](https://hub.docker.com/repository/docker/brpaz/go-api-sample), tagged wuth the short commit hash.

---

You can check the full pipeline definition [here](https://github.com/brpaz/go-api-sample/blob/master/.github/workflows/ci.yml)


## Build Docs

This workflow is triggered when there is a push to the `docs` folder of this project or it can also be run manually.

```yaml
on:
  push:
    branches:
      - master
    paths:
      - 'docs/**'
  workflow_dispatch:
```

It is reponsible for building the docs and deploy to GitHub Pages.


## Release

This workflow is run manually and creates a new release of the application. It pulls the latest Docker image that corresponds to the commit hash of the latest commit in master branch
and creates a new release on GitHub as well as tag the image in the Docker registry.

It uses [semantic-release](https://github.com/semantic-release/semantic-release) and [conventional-commits](https://www.conventionalcommits.org/en/v1.0.0/) to automatically
compute the next release version.

This would also be done in the CI pipeline when pushing to master branch if we want full Continuous Delivery workflow, but for this project and wanted to try manual releases.

<alert>
Right now, there is a limitation that if the latest commit in master, doesnt match a corresponding docker image in the registry, the build will fail.
This can happen, if you push documentation only updates which dont trigger a full image build.

We could change this to fetch the most recent tag by default.
</alert
