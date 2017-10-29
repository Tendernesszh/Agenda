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
	"fmt"

	"github.com/HinanawiTenshi/Agenda/entity"
	"github.com/spf13/cobra"
)

// quitmeetingCmd represents the quitmeeting command
var quitmeetingCmd = &cobra.Command{
	Use:   "quitmeeting",
	Short: "A brief description of your command",
	Long:  `quit the meeting current user particapated`,
	Run: func(cmd *cobra.Command, args []string) {
		AllMeeting := entity.GetMeetings()
		for _, meeting := range AllMeeting {
			for j, particapator := range meeting.Members {
				if curUser, _ := getCurUser(); particapator.Username == curUser {
					meeting.Members = append(meeting.Members[:j], meeting.Members[j+1:]...)
				}
			}
		}
		entity.UpdateMeeting(AllMeeting)
		fmt.Println("meeting updated")
	},
}

func init() {
	RootCmd.AddCommand(quitmeetingCmd)
	quitmeetingCmd.Flags().StringVarP(&_title, "title", "t", "", "meeting title")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
