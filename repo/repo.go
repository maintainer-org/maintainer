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

package repo

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

const (
	gitCmd                 string = "git"
	gitConfigArgs          string = "config"
	gitConfigGetArgs       string = "--get"
	gitConfigGetRemoteArgs string = "remote.origin.url"
)

var (
	gitRemotePattern = []string{
		`.*(?:[:\/])(?P<user>(?:-|\w|\.)*)\/(?P<project>(?:-|\w|\.)*)(?:\.git).*`,
		`.*\/(?P<user>(?:-|\w|\.)*)\/(?P<project>(?:-|\w|\.)*).*`,
	}

	errNameOrProjectNotExists = errors.New("Couldn't get the URL of this repository")

	singleton *Repository
)

// Repository is the type for the local repository.
type Repository struct {
	Owner string
	Name  string
}

// NewRepository returns a new Repository.
func NewRepository() (*Repository, error) {
	if singleton != nil {
		return singleton, nil
	}

	name, project, err := getNameAndRepoName()
	if err != nil {
		return nil, err
	}

	singleton = &Repository{
		Owner: name,
		Name:  project,
	}
	return singleton, nil
}

// String returns the information of a local repository.
func (r *Repository) String() string {
	return fmt.Sprintf("Local Repository Information: \n\tOwner: %s\n\tName: %s\n", r.Owner, r.Name)
}

// See https://github.com/skywinder/github-changelog-generator/blob/master/lib/github_changelog_generator/parser.rb#L312.
func getNameAndRepoName() (string, string, error) {
	cmd := exec.Command(gitCmd, gitConfigArgs, gitConfigGetArgs, gitConfigGetRemoteArgs)
	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}
	outputStr := string(output)

	// get the name and project.
	name, project, err := getNameAndRepoNameFromRemote(outputStr)
	if err != nil {
		return "", "", err
	}

	return name, project, nil
}

// getNameAndRepoName gets the name and project from local repository.
func getNameAndRepoNameFromRemote(remoteStr string) (string, string, error) {
	for _, regEx := range gitRemotePattern {
		paramsMap := getParams(regEx, remoteStr)
		name, ok1 := paramsMap["user"]
		project, ok2 := paramsMap["project"]
		if ok1 != true || ok2 != true {
			continue
		}
		return name, project, nil
	}
	return "", "", errNameOrProjectNotExists
}

// getParams get the params from regexp.
// See http://stackoverflow.com/questions/30483652/how-to-get-capturing-group-functionality-in-golang-regular-expressions.
func getParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}
