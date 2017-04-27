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

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gaocegege/maintainer/cmd"
	"github.com/spf13/cobra/doc"
)

const (
	outputDir     string = "docs"
	maintainerDoc string = "README.md"
)

// main creates docs for maintainer.
// go run ./gendoc/generate-doc.go
func main() {
	// Remove old docs.
	if err := os.RemoveAll(outputDir); err != nil {
		log.Fatalf("Error when delete old docs: %s", err)
	}
	// Create docs directory.
	if err := os.Mkdir(outputDir, 0775); err != nil {
		log.Fatalf("Error when create docs directory: %s", err)
	}

	// Create README.md.
	f, err := os.Create(fmt.Sprintf("%s/%s", outputDir, maintainerDoc))
	if err != nil {
		log.Fatalf("Error when create docs for maintainer: %s", err)
	}
	if err := doc.GenMarkdown(cmd.RootCmd, f); err != nil {
		log.Fatalf("Error when create docs for %s: %s", cmd.RootCmd.Name(), err)
	}

	// Create docs for subcommand.
	// Notice: There are no nested subcommands so no need to generate docs recursively.
	for _, clicmd := range cmd.RootCmd.Commands() {
		fileName := strings.Replace(clicmd.CommandPath(), " ", "_", -1)
		f, err := os.Create(fmt.Sprintf("%s/%s.md", outputDir, fileName))
		if err != nil {
			log.Fatalf("Error when create docs for %s: %s", clicmd.Name(), err)
		}
		if err := doc.GenMarkdown(clicmd, f); err != nil {
			log.Fatalf("Error when create docs for %s: %s", clicmd.Name(), err)
		}
	}
}
