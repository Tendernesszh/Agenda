package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
)

func init() {
  RootCmd.AddCommand(users)
}

var users = &cobra.Command{
  Use:   "users",
  Short: "show all the users",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("this is our users motherfucker")
  },
}
