// Copyright © 2017 Tendernesszh <dr.paper@live.com>
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

	"github.com/Tendernesszh/Agenda/util"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(loginCmd)
	// Initialize the flags
	registerCmd.Flags().StringVarP(&_username, "username", "u", "",
		"Specify the title of the meeting need to be created.")
	registerCmd.Flags().StringVarP(&_password, "password", "p",
		make([]string, 0), "Specify the members to attend the meeting.")

}

var loginCmd = &cobra.Command{
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
		if err := userloginArgsCheck(cmd); err != nil {
			fmt.Println(err)
			return
		}
    setCurUser(_username);
		fmt.Printf("[SUCCESS]User \"%v\" login\n", _username)
	},
}

func userloginArgsCheck(cmd *cobra.Command) error {
	users := util.GetUsers()
	//meetings := util.GetMeetings()

	// Check for the number of arguments
	if cmd.Flags().NFlag() != 2 {
		return argsError{invalidNArgs: true}
	}


	// Check for members that haven't registerred yet.

		for _, user := range users {
			if user.Username == _username {
				exist = true
        if user.Password == _password
        flag = true
				break
			}
		}

    if !flag {
      return argsError{permissionDeny: _username}
    }

		if !exist {
			return argsError{unknownUser: _username}
		}



	return nil
}