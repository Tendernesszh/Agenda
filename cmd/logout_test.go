package cmd

import(
  "io/ioutil"
  "testing"
)

func Testlogout(t *testing.T) {
  logoutCmd.Run(logoutCmd, make([]string, 0))
  name, _ := ioutil.ReadFile("curUser.txt")

  if name != nil {
    t.Errorf("fail")
  }
}
