---
title: 'Conventional Commits'
position: 3
category: 'Code quality'
fullscreen: false
---

This project follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification for commit messages.

The Conventional Commits specification is a lightweight convention on top of commit messages. 

It provides an easy set of rules for creating an explicit commit history; which makes it easier to write automated tools on top of. 

This convention follows [SemVer](https://semver.org/), by describing the features, fixes, and breaking changes made in commit messages. This allows to automate the release process by having tools like
[semantic-release](https://github.com/semantic-release/semantic-release) that inspects the commits to find which is the next version of the application.

The commit contains the following structural elements, to communicate intent to the consumers of your library:

* **fix**: a commit of the type fix patches a bug in your codebase (this correlates with PATCH in semantic versioning).
* **feat**: a commit of the type feat introduces a new feature to the codebase (this correlates with MINOR in semantic versioning).
* **BREAKING CHANGE**: a commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with MAJOR in semantic versioning). A BREAKING CHANGE can be part of commits of any type.

Types other than fix: and feat: are allowed, for example @commitlint/config-conventional (based on the the Angular convention) recommends **build**, **chore**, **ci**, **docs**, **style**, **refactor**, **perf**, **test**, and others.

A tool like [cz-cli](https://github.com/commitizen/cz-cli) can help by providing an interactive prompt to guide thought the creation of the commit messages.

<alert>
If you are using a feature branch I recommend, to just write normal commit messages and then "Squash and merge" into master with a proper Conventional Commits message.
</alert>
