package cmd

import (
	"io/ioutil"
	"testing"
)

func TestLogout(t *testing.T) {
	logoutCmd.Run(logoutCmd, make([]string, 0))

	name, _ := ioutil.ReadFile(CURUSER_PATH)

	if name != nil {
		t.Errorf("fail")
	}
}
