// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/viper"

	"github.com/maintainer-org/maintainer/pkg/config"
	"github.com/maintainer-org/maintainer/pkg/user"
	"github.com/maintainer-org/maintainer/pkg/util"
)

var (
	dailyFile *string
)

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Generate daily report for the user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString(config.Token)
		// Override token in CLI.
		if tokenValue != nil && *tokenValue != "" {
			log.Println("Found token in flag, override it.")
			token = *tokenValue
		}
		u := user.New(token)
		if username == nil || *username == "" {
			log.Panicf("Error when generate daily report: username is nil\n")
		}
		res, err := u.GetDailyToday(*username)
		if err != nil {
			log.Panicf("Error when generate daily report: %v\n", err)
		}
		f, err := util.OpenFile(*dailyFile)
		if err != nil {
			log.Panicf("Error when generate daily report: %v\n", err)
		}
		if _, err := f.WriteString(res); err != nil {
			log.Panicf("Error when generate daily report: %v\n", err)
		}
	},
}

func init() {
	userCmd.AddCommand(dailyCmd)

	dailyFile = dailyCmd.Flags().String("output", "./daily.md", "Output file")
}
