// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	entity "github.com/HinanawiTenshi/Agenda/entity"
	"github.com/spf13/cobra"
)

var titleToBeRemoved string

// removemeetingCmd represents the removemeeting command
var removemeetingCmd = &cobra.Command{
	Use:   "removemeeting",
	Short: "A brief description of your command",
	Long:  `Remove a meeting created by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		AllMeetings := entity.GetMeetings()

		for i, meeting := range AllMeetings {
			if CurUsername, _ := getCurUser(); meeting.Host == CurUsername && titleToBeRemoved == meeting.Title {
				AllMeetings = append(AllMeetings[:i], AllMeetings[i+1:]...)
			}
		}
		entity.UpdateMeeting(AllMeetings)
	},
}

func init() {
	RootCmd.AddCommand(removemeetingCmd)
	removemeetingCmd.PersistentFlags().StringVarP(&titleToBeRemoved, "title", "-t", "", "title of the meeting to be canceled")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removemeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removemeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
