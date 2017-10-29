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

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Command input error")
		} else if len(args) == 0 {
			RootCmd.Help()
		} else {
			flag := 0
			if args[0] == "createmeeting" {
				flag := 1
				createmeetingCmd.Help()
			}
			if args[0] == "meeting" {
				flag := 1
				meetingCmd.Help()
			}
			if args[0] == "modifymemeber" {
				flag := 1
				modifymemeberCmd.Help()
			}
			if args[0] == "register" {
				flag := 1
			  registerCmd.Help()
			}
			if args[0] == "login" {
				flag := 1
				loginCmd.Help()
			}
			if args[0] == "logout" {
				flag := 1
				logoutCmd.Help()
			}
			if args[0] == "users" {
				flag := 1
				usersCmd.Help()
			}
			if args[0] == "destroy" {
				flag := 1
				destroyCmd.Help()
			}
			if args[0] == "removemeeting" {
				flag := 1
				removemeetingCmd.Help()
			}
			if args[0] == "quitmeeting" {
				flag := 1
				quitmeetingCmd.Help()
			}
			if args[0] == "clearmeeting" {
				flag := 1
				clearmeetingCmd.Help()
			}
			if flag == 0 {
				fmt.Println("Command input error")
			}
		}
		fmt.Println("help called")
	},
}

func init() {
	RootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
