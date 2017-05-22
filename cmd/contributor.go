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
	"os/exec"
	"strings"

	"sort"

	"github.com/gaocegege/maintainer/config"
	"github.com/gaocegege/maintainer/util"
	"github.com/spf13/cobra"
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
	order *string
)

// Contributor is the type for contributor.
type Contributor struct {
	Name   string
	Commit int
}

// ContributorSlice is the type for slice of contributors.
type ContributorSlice []*Contributor

// Len is part of sort.Interface.
func (d ContributorSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d ContributorSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d ContributorSlice) Less(i, j int) bool {
	return d[i].Commit < d[j].Commit
}

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

	order = contributorCmd.PersistentFlags().String(config.Order, orderTime, "The order to compose Authors.md."+
		"(time, commit)")
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
	outputStr = outputStr[:len(outputStr)-1]
	contributors := strings.Split(outputStr, "\n")
	dict := make(map[string]int)
	for _, contributor := range contributors {
		contributor = strings.Trim(contributor, "'")
		if _, ok := dict[contributor]; ok != true {
			dict[contributor] = 1
		} else {
			dict[contributor] = dict[contributor] + 1
		}
	}
	return composeOrder(&dict)
}

// authorHeader returns the header to be written into AUTHORS.md.
func authorHeader() string {
	return "# Authors\n\n"
}

func composeOrder(data *map[string]int) error {
	contributors := make(ContributorSlice, 0, len(*data))
	for k, v := range *data {
		contributors = append(contributors, &Contributor{
			Name:   k,
			Commit: v,
		})
	}

	switch *order {
	case orderCommit:
		sort.Sort(sort.Reverse(contributors))
	}
	return writeToFile(contributors)
}

func orderByCommit(contributors ContributorSlice) error {
	return nil
}

func writeToFile(contributors ContributorSlice) error {
	// Output results to AUTHORS.md.
	f, err := util.OpenFile(authorFile)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(authorHeader()); err != nil {
		return err
	}
	for _, k := range contributors {
		if _, err := f.WriteString(k.Name); err != nil {
			return err
		}
		if _, err := f.WriteString("\n\n"); err != nil {
			return err
		}
	}
	// Write footer to the AUTHORS.md.
	if _, err := f.WriteString(Footer()); err != nil {
		return err
	}
	return nil
}
