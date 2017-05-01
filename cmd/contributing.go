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

package cmd

import (
	"log"

	"github.com/gaocegege/maintainer/contributing"
	"github.com/gaocegege/maintainer/util"

	"github.com/spf13/cobra"
)

const (
	contributingFile string = "CONTRIBUTING.md"
)

// contributingCmd represents the contributing command
var contributingCmd = &cobra.Command{
	Use:   "contributing",
	Short: "Generate CONTRIBUTING.md for your repository.",
	Long: `contributing subcommand will generate CONTRIBUTING.md for your repository, now
this file is a general version.

In the future, maintainer will detect languages and generate corresponding
documentation about programming language specific flow for contribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := contributingRun(); err != nil {
			log.Fatalf("Error when creating %s: %s\n", contributingFile, err)
			return
		}
		log.Printf("%s created successfully.\n", contributingFile)
	},
}

func init() {
	RootCmd.AddCommand(contributingCmd)
}

// contributingRun runs the real logic to generate CONTRIBUTING.md.
func contributingRun() error {
	// Output results to AUTHORS.md.
	f, err := util.OpenFile(contributingFile)
	if err != nil {
		return err
	}
	contributingText, err := contributing.GetContributing()
	if err != nil {
		return err
	}
	if _, err := f.WriteString(contributingText); err != nil {
		return err
	}
	// Write footer to the CONTRIBUTING.md.
	if _, err := f.WriteString(Footer()); err != nil {
		return err
	}
	return nil
}
