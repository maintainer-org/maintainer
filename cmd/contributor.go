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
	"os"
	"os/exec"
	"strings"

	"github.com/gaocegege/maintainer/util"
	"github.com/spf13/cobra"
)

const (
	gitCmd        string = "git"
	gitLogArgs    string = "log"
	gitFormatArgs string = "--format='%aN <%aE>'"
	authorFile    string = "AUTHORS.md"
)

// contributorCmd represents the contributor command
var contributorCmd = &cobra.Command{
	Use:   "contributor",
	Short: "generate AUTHORS.md for your repository.",
	Long: `contributor subcommand will generate AUTHORS.md. It gives the contributors more 
passion to contribute.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := contributorRun(); err != nil {
			log.Fatalf("Error when creating %s: %s\n", authorFile, err)
			return
		}
		log.Printf("%s created successfully.\n", authorFile)
	},
}

func init() {
	RootCmd.AddCommand(contributorCmd)
}

// contributorRun runs the real logic to generate AUTHORS.md.
func contributorRun() error {
	// git log --format='%aN <%aE>'.
	gitLogCmd := exec.Command(gitCmd, gitLogArgs, gitFormatArgs)
	output, err := gitLogCmd.Output()
	if err != nil {
		return err
	}

	// Parse output and remove duplicates.
	outputStr := string(output)
	contributors := strings.Split(outputStr, "\n")
	dict := make(map[string]int)
	for _, contributor := range contributors {
		contributor = strings.Trim(contributor, "'")
		if _, ok := dict[contributor]; ok != true {
			dict[contributor] = 1
		}
	}

	// Output results to AUTHORS.md.
	f, err := util.OpenFile(authorFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(authorHeader()); err != nil {
		return err
	}
	for k := range dict {
		if _, err := f.WriteString(k); err != nil {
			return err
		}
		if _, err := f.WriteString("\n"); err != nil {
			return err
		}
	}
	if _, err := f.WriteString(util.AuthorFooter()); err != nil {
		return err
	}
	return nil
}

// authorHeader returns the header to be written into AUTHORS.md.
func authorHeader() string {
	return "# Authors\n\n"
}
