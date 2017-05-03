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
	"fmt"
	"log"

	"github.com/gaocegege/maintainer/util"
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
	case util.LanguageGo:
		return c.getGoCodingStyle(), nil
	case util.LanguageJava:
		return c.getJavaCodingStyle(), nil
	case util.LanguageJavaScript:
		return c.getJavaScriptCodingStyle(), nil
	case util.LanguageScala:
		return c.getScalaCodingStyle(), nil
	case util.LanguageShell:
		return c.getShellCodingStyle(), nil
	case util.LanguageCpp:
		return c.getCppCodingStyle(), nil
	case util.LanguagePython:
		return c.getPythonCodingStyle(), nil
	case util.LanguagePHP:
		return c.getPHPCodingStyle(), nil
	case util.LanguageRuby:
		return c.getRubyCodingStyle(), nil
	case util.LanguageSwift:
		return c.getSwiftCodingStyle(), nil
	case util.LanguageR:
		return c.getRCodingStyle(), nil
	default:
		log.Printf("%s is not supported now to generate coding style guide.", language)
		return "", nil
	}
}

func (c *CodingStyleChooser) getCodingStyleTemplate() string {
	return `## Coding Style

See the [%s](%s) for details.

`
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

func (c *CodingStyleChooser) getJavaScriptCodingStyle() string {
	return `## Coding Style

The coding style suggested by the Golang community is used in this repository. See the 
[style doc](https://github.com/golang/go/wiki/CodeReviewComments) for details.

`
}

func (c *CodingStyleChooser) getScalaCodingStyle() string {
	return fmt.Sprintf(c.getCodingStyleTemplate(), "Scala style doc",
		"http://docs.scala-lang.org/style/")
}

func (c *CodingStyleChooser) getCppCodingStyle() string {
	return fmt.Sprintf(c.getCodingStyleTemplate(), "C++ style doc",
		"https://google.github.io/styleguide/cppguide.html")
}

func (c *CodingStyleChooser) getPythonCodingStyle() string {
	return fmt.Sprintf(c.getCodingStyleTemplate(), "Python style doc",
		"https://www.python.org/dev/peps/pep-0008/")
}

func (c *CodingStyleChooser) getShellCodingStyle() string {
	return fmt.Sprintf(c.getCodingStyleTemplate(), "Shell style doc",
		"https://google.github.io/styleguide/shell.xml")
}

func (c *CodingStyleChooser) getPHPCodingStyle() string {
	return fmt.Sprintf(c.getCodingStyleTemplate(), "PHP style doc",
		"http://www.php-fig.org/psr/psr-2/")
}

func (c *CodingStyleChooser) getRubyCodingStyle() string {
	return fmt.Sprintf(c.getCodingStyleTemplate(), "Ruby style doc",
		"https://github.com/bbatsov/ruby-style-guide")
}

func (c *CodingStyleChooser) getSwiftCodingStyle() string {
	return fmt.Sprintf(c.getCodingStyleTemplate(), "Swift style doc",
		"https://github.com/raywenderlich/swift-style-guide")
}

func (c *CodingStyleChooser) getRCodingStyle() string {
	return fmt.Sprintf(c.getCodingStyleTemplate(), "R style doc",
		"https://google.github.io/styleguide/Rguide.xml")
}
