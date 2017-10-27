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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const timeForm = "2006/01/02/15:04"

var cfgFile string

// Flags
var (
	_title      string
	_members    []string
	_member     string
	_starttime  string
	_endtime    string
	_addFlag    bool
	_removeFlag bool
)

// Current user's information(username)
const CurUserPath = "data/curUser.txt"

// Current User
var _username string

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

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "Agenda",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Agenda.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
