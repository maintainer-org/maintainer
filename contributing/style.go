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
)

const (
	languageGo   string = "Go"
	languageJava string = "Java"
)

// CodingStyleChooser is the chooser for coding style.
type CodingStyleChooser struct{}

// NewCodingStyleChooser creates a new CodingStyleChooser.
func NewCodingStyleChooser() *CodingStyleChooser {
	return &CodingStyleChooser{}
}

// GetCodingStyle gets the code style text according to language.
func (c *CodingStyleChooser) GetCodingStyle(language string) (string, error) {
	switch language {
	case languageGo:
		return c.getGoCodingStyle(), nil
	case languageJava:
		return c.getJavaCodingStyle(), nil
	default:
		log.Printf("%s is not supported now to generate coding style guide.", language)
		return "", nil
	}
}

func (c *CodingStyleChooser) getJavaCodingStyle() string {
	return `## Coding Style

The coding style in google is used in this repository. See the 
[Java style doc](https: //google.github.io/styleguide/javaguide.html) for details.

`
}

func (c *CodingStyleChooser) getGoCodingStyle() string {
	return `## Coding Style

The coding style suggested by the Golang community is used in this repository. See the 
[style doc](https://github.com/golang/go/wiki/CodeReviewComments) for details.

`
}
