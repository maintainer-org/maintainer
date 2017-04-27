# maintainer

[![Go Report Card](https://goreportcard.com/badge/github.com/gaocegege/maintainer)](https://goreportcard.com/report/github.com/gaocegege/maintainer)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gaocegege/maintainer)
[![Build Status](https://travis-ci.org/gaocegege/maintainer.svg?branch=master)](https://travis-ci.org/gaocegege/maintainer)
[![](https://img.shields.io/badge/docker-supported-blue.svg)](https://hub.docker.com/r/gaocegege/maintainer/)

Help you to be a qualified maintainer :)

Maintainer is a CLI app which helps you to generate AUTHORS, CHANGELOG.md, CONTRIBUTING.md and so on based on the repository in GitHub. It makes your repository more contributor-friendly.

[![asciicast](https://asciinema.org/a/117832.png)](https://asciinema.org/a/117832)

## Installation

```bash
$ go get github.com/gaocegege/maintainer
$ maintainer --help
```

## The Docker Way

[![Docker Pulls](https://img.shields.io/docker/pulls/gaocegege/maintainer.svg)](https://hub.docker.com/r/gaocegege/maintainer/)
[![](https://images.microbadger.com/badges/image/gaocegege/maintainer.svg)](https://microbadger.com/images/gaocegege/maintainer "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/gaocegege/maintainer.svg)](https://microbadger.com/images/gaocegege/maintainer "Get your own version badge on microbadger.com")

```bash
$ docker pull gaocegege/maintainer
$ docker run --rm -v $(pwd):/workdir gaocegege/maintainer:latest --help
```

## Config

### Location

`$HOME/.maintainer.yml`

### Configuration

Get the token from https://github.com/settings/tokens/new?description=Maintainer%20CLI%20token, you only need "repo" scope for private repositories.

```yaml
token: <Your Token for Maintainer CLI(40 digit)>
```

### Template

[./.maintainer.yml.template](./.maintainer.yml.template)

## Features

### Generate CHANGELOG.md

changelog subcommand will generate CHANGELOG.md for your repository, it is supported
via github_changelog_generator, so you need to install it before the subcommand is called.

See [here about how to install github_changelog_generator.](https://github.com/skywinder/github-changelog-generator#installation) In the future, maintainer will support install this dependency automatically.

### Generate CONTRIBUTING.md

contributing subcommand will generate CONTRIBUTING.md for your repository, now this file is a general version.

In the future, maintainer will detect languages and generate corresponding documentation about programming language specific flow for contribution.

### Generate AUTHORS

contributor subcommand will generate AUTHORS just like [moby/moby](https://github.com/moby/moby/blob/master/AUTHORS) does. It gives the contributors more passion to contribute.

### Recommend badges for you (In Future)

People :heart: badges. badge subcommand will recommend badges for you based on which languages the repository is written in.

## Docs

[./docs/README.md](./docs/README.md)
