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

	"github.com/JasonZang1005/Agenda/entity"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(usersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "A brief description of your command.....",
	Long:  `show all the users`,
	Run: func(cmd *cobra.Command, args []string) {
		curUser, _ := getCurUser()
		if curUser == "" {
			fmt.Println(argsError{permissionDeny: true}.Error())
			_errorLog.Println(argsError{permissionDeny: true}.Error())
			return
		}
		userList := entity.GetUsers()
		for _, user := range userList {
			fmt.Println("Username: ", user.Username)
			fmt.Println("Email: ", user.Email)
			fmt.Println("Phone: ", user.Phone)
		}
		_infoLog.Printf("[%v] Show users\n", curUser)
	},
}
