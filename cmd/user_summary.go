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
	"time"

	"github.com/maintainer-org/maintainer/pkg/config"
	"github.com/maintainer-org/maintainer/pkg/user"
	"github.com/maintainer-org/maintainer/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	summaryBegin, summaryEnd, summaryFile *string
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString(config.Token)
		// Override token in CLI.
		if tokenValue != nil && *tokenValue != "" {
			log.Println("Found token in flag, override it.")
			token = *tokenValue
		}
		u := user.New(token)
		if username == nil || *username == "" {
			log.Panicf("Error when generate summary: username is nil\n")
		}
		begin, err := time.Parse("2006-01-02", *summaryBegin)
		if err != nil {
			log.Panicf("Error when generate summary: %v\n", err)
		}
		end, err := time.Parse("2006-01-02", *summaryEnd)
		if err != nil {
			log.Panicf("Error when generate summary: %v\n", err)
		}
		res, err := u.GetSummary(*username, begin, end)
		if err != nil {
			log.Panicf("Error when generate summary: %v\n", err)
		}
		f, err := util.OpenFile(*summaryFile)
		if err != nil {
			log.Panicf("Error when generate summary: %v\n", err)
		}
		if _, err := f.WriteString(res); err != nil {
			log.Panicf("Error when generate summary: %v\n", err)
		}
	},
}

func init() {
	userCmd.AddCommand(summaryCmd)

	summaryBegin = summaryCmd.Flags().String("begin", "", "begin date of the summary (format 2011-11-11)")
	summaryEnd = summaryCmd.Flags().String("end", "", "end date of the summary (format 2011-11-11)")
	summaryFile = summaryCmd.Flags().String("output", "./summary.md", "Output file")
}
