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

	"github.com/gaocegege/maintainer/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	changelogGeneratorCmd string = "github_changelog_generator"
)

// changelogCmd represents the changelog command
var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Generate CHANGELOG.md for your repository.",
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
}

func changelogRun() error {
	token := viper.GetString(config.Token)
	// Override token in CLI.
	if *tokenValue != "" {
		log.Println("Found token in flag, override it.")
		token = *tokenValue
	}
	cmd := exec.Command(changelogGeneratorCmd, "-t", token)
	// Set STDERR and STDOUT to STDOUT of maintainer.
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
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
