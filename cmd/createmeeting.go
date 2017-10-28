// Copyright © 2017 HinanawiTenshi <dr.paper@live.com>
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
	"time"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(createmeetingCmd)

	// Initialize the flags
	createmeetingCmd.Flags().StringVarP(&_title, "title", "t", "",
		"Specify the title of the meeting need to be created.")
	createmeetingCmd.Flags().StringSliceVarP(&_members, "members", "m",
		make([]string, 20), "Specify the members to attend the meeting.")
	createmeetingCmd.Flags().StringVarP(&_starttime, "starttime", "s", "",
		"Specify the start time of the meeting in format yyyy/mm/dd/hh:mm")
	createmeetingCmd.Flags().StringVarP(&_endtime, "endtimeStr", "e", "",
		"Specify the end time of the meeting in format yyyy/mm/dd/hh:mm")
}

// createmeetingCmd represents the createmeeting command
var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "Create a meeting whose host is the current user.",
	Long: `Create a meeting with title, members, start time and end time.
The members must be users that have registerred, and if any members, including
you, is busy during the time, the meeting cannot be created.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 && len(args) == 0 {
			cmd.Help()
			return
		}
		if err := argsCheck(cmd); err != nil {
			fmt.Println(err)
			return
		}
		memberList := make([]SimpleUser, len(_members))
		for i := range memberList {
			memberList[i].Username = _members[i]
		}
		AddOneMeeting(
			Meeting{Title: _title, Members: memberList,
				Starttime: _starttime, Endtime: _endtime})
		fmt.Printf("[SUCCESS]Meeting \"%v\" created\n", _title)
	},
}

func argsCheck(cmd *cobra.Command) error {
	users := GetUsers()
	meetings := GetMeetings()

	// Check for the number of arguments
	if cmd.Flags().NFlag() != 4 {
		return argsError{invalidNArgs: true}
	}

	// Check for duplicated title
	for _, meeting := range meetings {
		if meeting.Title == _title {
			return argsError{duplicatedTitle: _title}
		}
	}

	// Check for members that haven't registerred yet.
	for _, member := range _members {
		exist := false
		for _, user := range users {
			if user.Username == member {
				exist = true
				break
			}
		}
		if !exist {
			return argsError{unknownUser: member}
		}
	}

	// Check for time
	st, errSt := time.Parse(timeForm, _starttime)
	et, errEt := time.Parse(timeForm, _endtime)
	if errSt != nil {
		return argsError{invalidArgs: "start time"}
	}
	if errEt != nil {
		return argsError{invalidArgs: "end Time"}
	}
	if st.After(et) || st.Equal(et) {
		return argsError{invalidArgs: "duration"}
	}

	// Check for busy members
	var busyMembers []string
	someoneBusy := false
	for _, user := range users {
		for _, meeting := range meetings {
			if meeting.HasUser(user.Username) {
				mSt, _ := time.Parse(timeForm, meeting.Starttime)
				mEt, _ := time.Parse(timeForm, meeting.Endtime)
				if !(mEt.Before(st) || mSt.After(et)) {
					busyMembers = append(busyMembers, user.Username)
					someoneBusy = true
					break
				}
			}
		}
	}
	if someoneBusy {
		return argsError{busyMembers: busyMembers}
	}

	return nil
}
