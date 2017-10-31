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

// modifymemeberCmd represents the modifymemeber command
var modifymemeberCmd = &cobra.Command{
	Use:   "modifymember",
	Short: "Add or remove members from your meeting",
	Long: `You can add or remove members corresponding to the meetings you
    created. You can not add a member to a meeting if the member is busy during
    the meeting. If the number of members of a meeting drops to 0, the meeting
    will be removed too.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check for arguments
		curUser, _ := getCurUser()
		if curUser == "" {
			fmt.Println(argsError{permissionDeny: true}.Error())
			_errorLog.Println(argsError{permissionDeny: true}.Error())
			return
		}
		if cmd.Flags().NFlag() == 0 && len(args) == 0 {
			cmd.Help()
			return
		}
		if cmd.Flags().NFlag() != 3 {
			fmt.Println(argsError{invalidNArgs: true}.Error())
			_errorLog.Println(argsError{invalidNArgs: true}.Error())
			return
		}
		meetings := entity.GetMeetings()
		users := entity.GetUsers()
		validTitle := false
		for _, meeting := range meetings {
			if meeting.Title == _title {
				validTitle = true
				break
			}
		}
		if !validTitle {
			fmt.Println(argsError{unknownTitle: _title}.Error())
			_errorLog.Println(argsError{unknownTitle: _title}.Error())
			return
		}
		meeting := entity.GetMeeting(_title)
		if meeting.Host != curUser {
			fmt.Println(argsError{permissionDeny: true}.Error())
			_errorLog.Println(argsError{permissionDeny: true}.Error())
			return
		}
		busy := make([]string, 0)
		for _, member := range _members {
			exist := false
			for _, user := range users {
				if user.Username == member {
					exist = true
					break
				}
			}
			st, _ := time.Parse(TIME_FORM, meeting.Starttime)
			et, _ := time.Parse(TIME_FORM, meeting.Endtime)
			for _, m := range meetings {
				if m.Title == _title {
					continue
				}
				if m.HasUser(member) {
					curSt, _ := time.Parse(TIME_FORM, m.Starttime)
					curEt, _ := time.Parse(TIME_FORM, m.Endtime)
					if !(curEt.Before(st) || curEt.Equal(st) ||
						curSt.After(et) || curSt.Equal(et)) {
						busy = append(busy, member)
						break
					}
				}
			}
			if !exist {
				fmt.Println(argsError{unknownUser: member}.Error())
				_errorLog.Println(argsError{unknownUser: member}.Error())
				return
			}
		}
		if len(busy) != 0 {
			fmt.Println(argsError{busyMembers: busy}.Error())
			_errorLog.Println(argsError{busyMembers: busy}.Error())
			return
		}

		// Modified the members
		for i, meeting := range meetings {
			if meeting.Title == _title {
				if _addFlag {
					for _, newMember := range _members {
						for _, oldMember := range meeting.Members {
							if oldMember.Username == newMember {
								fmt.Println(argsError{duplicatedUser: newMember}.Error())
								_errorLog.Println(argsError{duplicatedUser: newMember}.Error())
								return
							}
						}
						meetings[i].Members = append(meeting.Members,
							entity.SimpleUser{Username: newMember})
					}
				} else if _removeFlag {
					newMembers := make([]entity.SimpleUser, 0)
					for _, oldMember := range meeting.Members {
						needRemove := false
						for _, rmMember := range _members {
							if rmMember == oldMember.Username {
								needRemove = true
							}
						}
						if !needRemove {
							newMembers = append(newMembers, oldMember)
						}
					}
					meetings[i].Members = newMembers
				}
				entity.UpdateMeeting(meetings)
				break
			}
		}
		_infoLog.Println("["+curUser+"] Modify \"%v\"'s members\n", _title)
	},
}

func init() {
	RootCmd.AddCommand(modifymemeberCmd)

	modifymemeberCmd.Flags().BoolVarP(&_addFlag, "add", "a", false,
		"Add members to the meeting")
	modifymemeberCmd.Flags().BoolVarP(&_removeFlag, "remove", "r", false,
		"Remove members from the meeting")
	modifymemeberCmd.Flags().StringVarP(&_title, "title", "t", "",
		"Specify the title of the meeting")
	modifymemeberCmd.Flags().StringSliceVarP(&_members, "members", "m",
		make([]string, 0), "Specify the members to work with")
}
