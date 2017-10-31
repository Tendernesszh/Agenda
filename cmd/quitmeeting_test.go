package cmd

import (
	"testing"

	"github.com/JasonZang1005/Agenda/entity"
)

func TestQuitmeeting(t *testing.T) {
	quitmeetingCmd.Run(quitmeetingCmd, make([]string, 0))
	meeting := entity.GetMeeting(_title)
	curUsername, _ := getCurUser()
	for _, member := range meeting.Members {
		if member.Username == curUsername {
			t.Errorf("user is still in the meeting")
		}
	}
}
