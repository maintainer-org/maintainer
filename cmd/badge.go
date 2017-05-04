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

	"github.com/spf13/cobra"

	"github.com/gaocegege/maintainer/badge"
)

// badgeCmd represents the badge command
var badgeCmd = &cobra.Command{
	Use:   "badge",
	Short: "Recommend badges according to your repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := badgeRun(); err != nil {
			log.Fatalf("Error when recommending badges: %s\n", err)
			return
		}
	},
}

func badgeRun() error {
	recommender, err := badge.NewRecommender()
	if err != nil {
		return err
	}
	if err = recommender.Recommend(); err != nil {
		return err
	}
	return nil
}

func init() {
	RootCmd.AddCommand(badgeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// badgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// badgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
