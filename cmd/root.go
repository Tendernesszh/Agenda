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
	"io/ioutil"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const TIME_FORM = "2006/01/02/15:04"

var cfgFile string

// Flags
var (
	_username   string
	_password   string
	_email      string
	_phone      string
	_title      string
	_members    []string
	_member     string
	_starttime  string
	_endtime    string
	_addFlag    bool
	_removeFlag bool
)

// Current user's information(username)
const CURUSER_PATH = "data/curUser.txt"

// ERROR
type argsError struct {
	invalidNArgs    bool
	invalidArgs     string
	duplicatedTitle string
	duplicatedUser  string
	unknownTitle    string
	unknownUser     string
	busyMembers     []string
	permissionDeny  bool
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
			fmt.Sprintf("[ERROR]Meeting \"%v\" already existed", e.duplicatedTitle)
	}
	if e.duplicatedUser != "" {
		result +=
			fmt.Sprintf("[ERROR]Member \"%v\" already existed", e.duplicatedUser)
	}
	if e.unknownTitle != "" {
		result +=
			fmt.Sprintf("[ERROR]Meeting \"%v\" not found", e.unknownTitle)
	}
	if e.unknownUser != "" {
		result += fmt.Sprintf("[ERROR]Unknown user %v", e.unknownUser)
	}
	if len(e.busyMembers) > 0 {
		busy := "[ERROR]The following members are busy during the time\n"
		for busyMem := range e.busyMembers {
			if busyMem == len(e.busyMembers)-1 {
				busy += e.busyMembers[busyMem]
			} else {
				busy += e.busyMembers[busyMem] + " "
			}
		}
		result += busy
	}
	if e.permissionDeny {
		result += "[ERROR]Permission Denied"
	}
	return result
}

func timeIntervalCheck() error {
	st, errSt := time.Parse(TIME_FORM, _starttime)
	et, errEt := time.Parse(TIME_FORM, _endtime)
	if errSt != nil {
		return argsError{invalidArgs: "start time"}
	}
	if errEt != nil {
		return argsError{invalidArgs: "end Time"}
	}
	if st.After(et) || st.Equal(et) {
		return argsError{invalidArgs: "duration"}
	}
	return nil
}

func getCurUser() (string, error) {
	name, err := ioutil.ReadFile(CURUSER_PATH)
	if err != nil {
		return "", err
	}
	return string(name), nil
}

func setCurUser(username string) error {
	if username == "" {
		os.Truncate(CURUSER_PATH, 0)
		return nil
	}
	file, err := os.OpenFile(CURUSER_PATH, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	file.Write([]byte(username))
	return nil
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "Agenda",
	Short: "A useful meetings management tool",
	Long: `By using Agenda, you can create your own account and do easy
meetings managements among your partners. Make sure they are all
registerred here.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".Agenda" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".Agenda")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
