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

	"github.com/spf13/cobra"
)

const (
	gitHubTocCmd string = "gh-md-toc"
	readmeFile   string = "readme.md"
)

// readmeTocCmd represents the readme command
var readmeTocCmd = &cobra.Command{
	Use:   "toc",
	Short: "Generate TOC for your repository",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := readmeTocRun(); err != nil {
			log.Fatalf("Error when generating TOC: %s\n", err)
			return
		}
		log.Printf("TOC generated successfully.\n")
	},
}

func init() {
	readmeCmd.AddCommand(readmeTocCmd)
}

func readmeTocRun() error {
	cmd := exec.Command(gitHubTocCmd, readmeFile)
	// Set STDERR and STDOUT to STDOUT of maintainer.
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
