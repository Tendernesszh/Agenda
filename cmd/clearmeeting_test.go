package cmd

import(
  "io/ioutil"
  "testing"
)

func Testclearmeeting(t *testing.T) {
  clearmeetingCmd.Run(clearmeetingCmd, make([]string, 0))
  information, _ := ioutil.ReadFile("meeting.json")

  if information != nil {
    t.Errorf("fail")
  }
}
