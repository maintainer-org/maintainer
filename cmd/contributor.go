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
	"context"
	"errors"
	"log"

	"golang.org/x/oauth2"

	"fmt"

	"github.com/gaocegege/maintainer/config"
	"github.com/gaocegege/maintainer/repo"
	"github.com/gaocegege/maintainer/util"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	gitCmd        string = "git"
	gitLogArgs    string = "log"
	gitFormatArgs string = "--format='%aN <%aE>'"
	authorFile    string = "AUTHORS.md"

	orderTime   string = "time"
	orderCommit string = "commit"
)

var (
	contributorPattern = []string{
		"(?P<user>.*) <(?P<email>.*@.*)>",
	}
	errNameOrEmailNotExists = errors.New("Couldn't get the name or email of one contributor")

	order *string
)

// contributorCmd represents the contributor command
var contributorCmd = &cobra.Command{
	Use:   "contributor",
	Short: "Generate AUTHORS.md for your repository.",
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

	order = contributorCmd.PersistentFlags().String(config.Order, orderCommit, "The order to compose Authors.md."+
		"(commit)")
}

// contributorRun runs the real logic to generate AUTHORS.md.
func contributorRun() error {
	repo, err := repo.NewRepository()
	if err != nil {
		log.Panicf("Error when read the information from local repository: %s\n", err)
	}

	token := viper.GetString(config.Token)
	// Override token in CLI.
	if *tokenValue != "" {
		log.Println("Found token in flag, override it.")
		token = *tokenValue
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	contributors := []*github.Contributor{}
	i := 1
	for {
		contributorsBuf, _, err := client.Repositories.ListContributors(repo.Owner, repo.Name, &github.ListContributorsOptions{
			// See https://developer.github.com/v3/repos/#list-contributors
			// Anon: "true",
			ListOptions: github.ListOptions{
				Page:    i,
				PerPage: 100,
			},
		})
		if err != nil {
			return err
		}
		if len(contributorsBuf) == 0 {
			break
		}
		contributors = append(contributors, contributorsBuf...)
		i = i + 1
	}

	if err := composeByOrder(contributors); err != nil {
		return err
	}
	return nil
}

// authorHeader returns the header to be written into AUTHORS.md.
func authorHeader() string {
	return "# Authors\n\nThis list is sorted by the number of commits per contributor in descending order.\n\n"
}

func composeByOrder(contributors []*github.Contributor) error {
	return writeToFile(contributors)
}

func writeToFile(contributors []*github.Contributor) error {
	// Output results to AUTHORS.md.
	f, err := util.OpenFile(authorFile)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(authorHeader()); err != nil {
		return err
	}
	for _, contributor := range contributors {
		if _, err := f.WriteString(fmt.Sprintf("* [@%s](%s)", *contributor.Login, *contributor.HTMLURL)); err != nil {
			return err
		}
		if _, err := f.WriteString("\n"); err != nil {
			return err
		}
	}
	// Write footer to the AUTHORS.md.
	if _, err := f.WriteString(Footer()); err != nil {
		return err
	}
	return nil
}
