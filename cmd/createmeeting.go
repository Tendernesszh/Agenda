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

	"github.com/HinanawiTenshi/Agenda/util"
	"github.com/spf13/cobra"
)

const timeForm = "2006/01/02/15:04"

// Flags
var (
	title     string
	members   []string
	starttime string
	endtime   string
)

// ERROR
type argsError struct {
	invalidNArgs    bool
	invalidArgs     string
	duplicatedTitle string
	unknownUser     string
	busyMembers     []string
}

func (e argsError) Error() string {
	var result string
	if e.invalidNArgs {
		result += "[ERROR]Arguments not fit"
	}
	if e.invalidArgs != "" {
		result += fmt.Sprintf("[ERROR]Invalid %v", e.invalidArgs)
	}
	if e.duplicatedTitle != "" {
		result +=
			fmt.Sprintf("[ERROR]\"%v\" already existed", e.duplicatedTitle)
	}
	if e.unknownUser != "" {
		result += fmt.Sprintf("[ERROR]Unknown user %v", e.unknownUser)
	}
	if len(e.busyMembers) > 0 {
		busy := `[ERROR]The following members are busy during
            the time\n`
		for busyMem := range e.busyMembers {
			if busyMem == len(e.busyMembers)-1 {
				busy += e.busyMembers[busyMem]
			} else {
				busy += e.busyMembers[busyMem] + " "
			}
		}
		result += busy
	}
	return result
}

func init() {
	RootCmd.AddCommand(createmeetingCmd)

	// Initialize the flags
	createmeetingCmd.Flags().StringVarP(&title, "title", "t", "",
		"Specify the title of the meeting need to be created.")
	createmeetingCmd.Flags().StringSliceVarP(&members, "members", "m",
		make([]string, 20), "Specify the members to attend the meeting.")
	createmeetingCmd.Flags().StringVarP(&starttime, "starttime", "s", "",
		"Specify the start time of the meeting in format yyyy/mm/dd/hh:mm")
	createmeetingCmd.Flags().StringVarP(&endtime, "endtimeStr", "e", "",
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
		memberList := make([]util.SimpleUser, len(members))
		for i := range memberList {
			memberList[i].Username = members[i]
		}
		util.AddOneMeeting(
			util.Meeting{Title: title, Members: memberList,
				Starttime: starttime, Endtime: endtime})
		fmt.Printf("[SUCCESS]Meeting \"%v\" created\n", title)
	},
}

func argsCheck(cmd *cobra.Command) error {
	// Check for the number of arguments
	if cmd.Flags().NFlag() != 4 {
		return argsError{invalidNArgs: true}
	}

	// Check for duplicated title
	meetings := util.GetMeetings()
	for _, meeting := range meetings {
		if meeting.Title == title {
			return argsError{duplicatedTitle: title}
		}
	}

	// Check for members that haven't registerred yet.
	users := util.GetUsers()
	for _, member := range members {
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
	st, errSt := time.Parse(timeForm, starttime)
	et, errEt := time.Parse(timeForm, endtime)
	if errSt != nil {
		return argsError{invalidArgs: "start time"}
	}
	if errEt != nil {
		return argsError{invalidArgs: "end Time"}
	}
	if st.After(et) || st.Equal(et) {
		return argsError{invalidArgs: "duration"}
	}

	// TODO: Check for busy members

	return nil
}
