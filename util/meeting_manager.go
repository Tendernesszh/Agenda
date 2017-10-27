package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type Meeting struct {
	Title     string       `json:"title"`
	Members   []SimpleUser `json:"members"`
	Starttime string       `json:"start_time"`
	Endtime   string       `json:"end_time"`
}

var (
	MeetingPath string = "data/meetings.json"
	meetingList []Meeting
)

func init() {
	file, err := os.OpenFile(MeetingPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonDecoder := json.NewDecoder(file)
	for jsonDecoder.More() {
		var curMeeting Meeting
		err := jsonDecoder.Decode(&curMeeting)
		if err != nil {
			fmt.Println(err)
			return
		}
		meetingList = append(meetingList, curMeeting)
	}
}

func AddOneMeeting(m Meeting) {
	file, err := os.OpenFile(MeetingPath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonEncoder := json.NewEncoder(file)
	encodeErr := jsonEncoder.Encode(&m)
	if encodeErr != nil {
		fmt.Println(encodeErr)
		return
	}
}

func GetMeetings() []Meeting {
	return meetingList
}

func (m *Meeting) HasUser(username string) bool {
	for _, member := range m.Members {
		if member.Username == username {
			return true
		}
	}
	return false
}
