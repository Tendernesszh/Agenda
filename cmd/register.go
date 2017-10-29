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

	"github.com/HinanawiTenshi/Agenda/entity"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(registerCmd)

	// Initialize the flags
	registerCmd.Flags().StringVarP(&_username, "username", "u", "",
		"Specify the username need to be created.")
	registerCmd.Flags().StringVarP(&_password, "password", "p",
		"", "Specify the password of the user.")
	registerCmd.Flags().StringVarP(&_email, "email", "e", "",
		"Specify the email of the user")
	registerCmd.Flags().StringVarP(&_phone, "phone", "ph", "",
		"Specify the phone of the user")
}

var registerCmd = &cobra.Command{
	Use:   "Register",
	Short: "Register a user account.",
	Long:  `Register a user account with username, password, email and phone.`,

	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 && len(args) == 0 {
			cmd.Help()
			return
		}
		if err := userArgsCheck(cmd); err != nil {
			fmt.Println(err)
			return
		}

		entity.AddOneUser(
			entity.User{Username: _username, Password: _password,
				Email: _email, Phone: _phone})
		fmt.Printf("[SUCCESS]User \"%v\" created\n", _username)
	},
}

func userArgsCheck(cmd *cobra.Command) error {
	users := entity.GetUsers()
	//meetings := entity.GetMeetings()

	// Check for the number of arguments
	if cmd.Flags().NFlag() != 4 {
		return argsError{invalidNArgs: true}
	}

	// Check for members that haven't registerred yet.
	exist := false
	for _, user := range users {
		if user.Username == _username {
			exist = true
			break
		}
	}
	if exist {
		return argsError{duplicatedUser: _username}
	}

	return nil
}
