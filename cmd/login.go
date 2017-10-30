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

	"github.com/JasonZang1005/Agenda/entity"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(loginCmd)
	// Initialize the flags
	loginCmd.Flags().StringVarP(&_username, "username", "u", "",
		"Specify the username.")
	loginCmd.Flags().StringVarP(&_password, "password", "p",
		"", "Specify the password.")

}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login account.",
	Long:  `login account. Please enter the correct password.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 && len(args) == 0 {
			cmd.Help()
			return
		}
		if err := userloginArgsCheck(cmd); err != nil {
			fmt.Println(err)
			return
		}
		setCurUser(_username)
		fmt.Printf("[SUCCESS]User \"%v\" login\n", _username)
	},
}

func userloginArgsCheck(cmd *cobra.Command) error {
	users := entity.GetUsers()
	//meetings := entity.GetMeetings()

	// Check for the number of arguments
	if cmd.Flags().NFlag() != 2 {
		return argsError{invalidNArgs: true}
	}

	// Check for members that haven't registerred yet.
	flag := false
	exist := false
	for _, user := range users {
		if user.Username == _username {
			exist = true
			if user.Password == _password {
				flag = true
				break
			}
		}
	}

	if !flag {
		return argsError{permissionDeny: true}
	}

	if !exist {
		return argsError{unknownUser: _username}
	}

	return nil
}
