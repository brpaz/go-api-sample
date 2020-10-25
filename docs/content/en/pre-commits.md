---
title: 'Pre-Commit hooks'
position: 2
category: 'Code quality'
fullscreen: false
---

Pre-commit hooks are a great way to reduce the chances of pushing broken code to the Git Repository.

[Pre-commit](https://pre-commit.com/) is a framework for managing and maintaining multi-language pre-commit hooks.

With a simple configuration file you can define which commands to run before certain Git operations like commit or push.

This project comes with a `pre-commit-config.yaml` file that runs the format, lint and unit tests before each commit as well some default pre-commit tasks like checking for large files or removing trailing white space.

```yaml
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v2.0.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-added-large-files
  - repo: local
    hooks:
      - id: format
        name: Format code
        entry: make fmt
        language: system
        types: [go]

      - id: lint
        name: Lint
        entry: make lint
        language: system
        pass_filenames: false

      - id: unittests
        name: Unit tests
        entry: make test
        language: system
        pass_filenames: false
        stages: ["push"]
```

For the hooks to work, you must first install pre-commit globally on your machine and then install the hooks on the repository. 

`make setup-env` should take care of this for you, by installing pre-commit using `pip`.
 
 If you want to install pre-commit in any other way, please check the [install](https://pre-commit.com/#install) guide on pre-commit website.



