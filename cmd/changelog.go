// Copyright © 2017 Ce Gao <ce.gao@outlook.com>
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
	changelogGeneratorCmd string = "github_changelog_generator"
)

// changelogCmd represents the changelog command
var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "generate changelog for your repository.",
	Long: `changelog subcommand will generate CHANGELOG.md for your repository, it is supported
via github_changelog_generator, so you need to install it before the subcommand is called.

In the future, maintainer will support install this dependency automatically.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := changelogRun(); err != nil {
			log.Fatalf("Error when creating changelog: %s\n", err)
			return
		}
		log.Println("changelog created successfully.")
	},
}

func init() {
	RootCmd.AddCommand(changelogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changelogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changelogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func changelogRun() error {
	cmd := exec.Command(changelogGeneratorCmd)
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// checkRequirements checks whether changelogGeneratorCmd is installed.
func checkRequirements() error {
	// TODO
	return nil
}
