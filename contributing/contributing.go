// Copyright © 2017 Maintainer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package contributing

import (
	"log"

	"fmt"

	"github.com/gaocegege/maintainer/repo"
	"github.com/google/go-github/github"
)

const (
	// startText is the text at the start.
	startText = `# How to contribute

This document outlines some of the conventions on development workflow, commit message formatting, contact points and other
resources to make it easier to get your contribution accepted.

## Getting started

- Fork the repository on GitHub.
- Read the README.md for build instructions.
- Play with the project, submit bugs, submit patches!

## Contribution flow

This is a rough outline of what a contributor's workflow looks like:

- Create a topic branch from where you want to base your work. This is usually master.
- Make commits of logical units and add test case if the change fixes a bug or adds new functionality.
- Run tests and make sure all the tests are passed.
- Make sure your commit messages are in the proper format (see below).
- Push your changes to a topic branch in your fork of the repository.
- Submit a pull request to the repo.

Thanks for your contributions!

`

	// endText is the text at the end.
	endText string = `### Format of the Commit Message

We follow a rough convention for commit messages that is designed to answer two
questions: what changed and why. The subject line should feature the what and
the body of the commit should describe the why.

<pre><code>
store/localstore: add comment for variable declaration.

Improve documentation.
</code></pre>

The format can be described more formally as follows:

<pre><code>
subsystem: what changed
BLANK LINE
why this change was made
BLANK LINE
footer(optional)
</code></pre>

The first line is the subject and should be no longer than 70 characters, the
second line is always blank, and other lines should be wrapped at 80 characters.
This allows the message to be easier to read on GitHub as well as in various
git tools.

If the change affects more than one subsystem, you can use comma to separate them like util/codec,util/types:.

If the change affects many subsystems, you can use * instead, like *:.

For the why part, if no specific reason for the change,
you can use one of some generic reasons like "Improve documentation.",
"Improve performance.", "Improve robustness.", "Improve test coverage."

`
)

// GetContributing gets the contributing text.
func GetContributing() (string, error) {
	codingStyleText, err := getCodingStyle()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s%s", startText, codingStyleText, endText), nil
}

// getCodingStyle gets the text about coding style.
func getCodingStyle() (string, error) {
	repo, err := repo.NewRepository()
	if err != nil {
		log.Panicf("Error when read the information from local repository: %s\n", err)
	}
	client := github.NewClient(nil)
	remoteRepo, _, err := client.Repositories.Get(repo.Owner, repo.Name)
	if err != nil {
		return "", err
	}

	chooser := NewCodingStyleChooser()
	text, err := chooser.GetCodingStyle(*remoteRepo.Language)
	if err != nil {
		return "", err
	}
	return text, nil
}
