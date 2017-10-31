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
	"time"

	"github.com/HinanawiTenshi/Agenda/entity"
	"github.com/spf13/cobra"
)

// meetingCmd represents the meeting command
var meetingCmd = &cobra.Command{
	Use:   "meeting",
	Short: "Query meetings of a specific time interval",
	Long: `You can use this command to query all meetings of a specific time
    interval, including the meetings you held and you participated`,
	Run: func(cmd *cobra.Command, args []string) {
		curUser, _ := getCurUser()
		meetings := entity.GetMeetings()
		if curUser == "" {
			fmt.Println(argsError{permissionDeny: true}.Error())
			_errorLog.Println(argsError{permissionDeny: true}.Error())
			return
		}
		if cmd.Flags().NFlag() == 0 && len(args) == 0 {
			fmt.Printf("Title\tStart time\tEnd time\tHost\tParticipants\n")
			for _, meeting := range meetings {
				entity.PrintOneMeeting(meeting)
			}
			_infoLog.Printf("[" + curUser + "] Show all meetings\n")
			return
		}
		if timeErr := timeIntervalCheck(); timeErr != nil {
			fmt.Println(timeErr)
			return
		}
		st, _ := time.Parse(TIME_FORM, _starttime)
		et, _ := time.Parse(TIME_FORM, _endtime)
		noMeeting := true
		for _, meeting := range meetings {
			cSt, _ := time.Parse(TIME_FORM, meeting.Starttime)
			cEt, _ := time.Parse(TIME_FORM, meeting.Endtime)
			if !(et.Before(cSt) || st.After(cEt) &&
				meeting.HasUser(curUser)) {
				if noMeeting {
					fmt.Printf("Title\tStart time\tEnd time\tHost\tParticipants\n")
					noMeeting = false
				}
				entity.PrintOneMeeting(meeting)
			}
		}
		if noMeeting {
			fmt.Println("No meetings here")
		}
		_infoLog.Printf("["+curUser+"] Show meetings within %v and %v\n",
			_starttime, _endtime)
	},
}

func init() {
	RootCmd.AddCommand(meetingCmd)

	//Initialize the flags
	meetingCmd.Flags().StringVarP(&_starttime, "starttime", "s", "",
		"The start time of the meetings")
	meetingCmd.Flags().StringVarP(&_endtime, "endtime", "e", "",
		"The end time of the meetings")
}
