<h1 align="center">
        <br>
        <img width="200" src="docs/img/logo.png" alt="maintainer">
        <br>
        <h4 align="center">Help you to be a qualified maintainer.</h4>
        <br>
</h1>


[![Go Report Card](https://goreportcard.com/badge/github.com/gaocegege/maintainer)](https://goreportcard.com/report/github.com/gaocegege/maintainer)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gaocegege/maintainer)
[![Build Status](https://travis-ci.org/gaocegege/maintainer.svg?branch=master)](https://travis-ci.org/gaocegege/maintainer)
[![](https://img.shields.io/badge/docker-supported-blue.svg)](https://hub.docker.com/r/gaocegege/maintainer/)
[![Libraries.io for GitHub](https://img.shields.io/librariesio/github/gaocegege/maintainer.svg)](https://libraries.io/github/gaocegege/maintainer)
[![Gitter](https://badges.gitter.im/gaocegege/maintainer.svg)](https://gitter.im/gaocegege/maintainer?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Maintainer is a CLI app which helps you to generate AUTHORS, CHANGELOG.md, CONTRIBUTING.md and so on based on the repository in GitHub. It makes your repository more contributor-friendly.

## Installation

```text
$ go get github.com/gaocegege/maintainer
$ maintainer --help
maintainer is a CLI app to help you to generate AUTHORS, CONTRIBUTING.md, CHANGELOG.md and so on.

Usage:
  maintainer [command]

Available Commands:
  changelog    Generate CHANGELOG.md for your repository.
  contributing Generate CONTRIBUTING.md for your repository.
  contributor  Generate AUTHORS.md for your repository.
  readme       Generate README.md related text in your repository.

Flags:
  --config string   config file (default is $HOME/.maintainer.yaml)
  -h, --help            help for maintainer

Use "maintainer [command] --help" for more information about a command.
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

There is only one configuration now: the token in GitHub. The token can be created from https://github.com/settings/tokens/new?description=Maintainer%20CLI%20token, you only need "repo" scope for private repositories.

It is used in changelog subcommand, so there are three ways to config the token:

* Config `token` field in `$HOME/.maintainer.yml`. [./.maintainer.yml.template](./.maintainer.yml.template) is a template.
* Or Set environment variable `MAINTAINER_TOKEN`.
* Or Set the flag in changelog subcommand: `maintainer changelog --token <token>`. If you set it in command, it will override the configuration in config file and the environment variable.

## Features Preview

[![asciicast](https://asciinema.org/a/117832.png)](https://asciinema.org/a/117832)

### Generate CHANGELOG.md

changelog subcommand will generate CHANGELOG.md for your repository, it is supported
via github_changelog_generator, so you need to install it before the subcommand is called.

See [here about how to install github_changelog_generator.](https://github.com/skywinder/github-changelog-generator#installation) In the future, maintainer will support install this dependency automatically.

See [CHANGELOG.md](./CHANGELOG.md) as an example.

### Generate CONTRIBUTING.md

contributing subcommand will generate CONTRIBUTING.md for your repository, now this file is a general version.

In the future, maintainer will detect languages and generate corresponding documentation about programming language specific flow for contribution.

See [CONTRIBUTING.md](./CONTRIBUTING.md) as an example.

### Generate AUTHORS.md

contributor subcommand will generate AUTHORS just like [moby/moby](https://github.com/moby/moby/blob/master/AUTHORS) does. It gives the contributors more passion to contribute.

See [AUTHORS.md](./AUTHORS.md) as an example.

### Generate README.md related text

readme subcommand contains some subcommands:

* `toc` subcommand, which could generate table of contents(TOC) for README.md. It is supported with [gh-md-toc](https://github.com/ekalinin/github-markdown-toc), so you need to install it before the subcommand is called.

### Recommend badges for you (Soon)

People :heart: badges. badge subcommand will recommend badges for you based on which languages the repository is written in.

## CLI references

[references/README.md](./references/README.md)

## Development

### Run from source code

```bash
go run main.go <subcommand>
```

### Build Docker image

Maintainer requires two steps to build a Docker image.

1. `scripts/build-for-alphine.sh` builds `maintainer` in a Docker container which from `golang:1.8-alpine`. It mounts maintainer directory into the container so the `maintainer` built from code will visiable in host.
1. `docker build -t maintainer .` builds real image from `Dockerfile`. It simply copys binary `maintainer` into the image and install some dependencies such as git and github_changelog_generator.

This way is inspired by [caicloud/cyclone](https://github.com/caicloud/cyclone). It could reduce the size of image significantly.

## Acknowledgments

* Thanks [github.com/spf13/cobra](https://github.com/spf13/cobra) for its powerful CLI generator.
* Thanks [github_changelog_generator](https://github.com/skywinder/github-changelog-generator) which is the source of the idea.
* Thanks [gophers](https://github.com/egonelbre/gophers) and [johan/1007813](https://gist.github.com/johan/1007813) for the awesome gophers and octocat.
