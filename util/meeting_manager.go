package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type Meeting struct {
	Title     string       `json:"title"`
	Host      string       `json:"host"`
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

func UpdateMeeting(meetings []Meeting) {
	file, err := os.OpenFile(MeetingPath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Truncate(MeetingPath, 0)
	defer file.Close()

	jsonEncoder := json.NewEncoder(file)
	for _, m := range meetings {
		encodeErr := jsonEncoder.Encode(&m)
		if encodeErr != nil {
			fmt.Println(encodeErr)
			return
		}
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

func GetMeeting(title string) Meeting {
	for _, meeting := range meetingList {
		if meeting.Title == title {
			return meeting
		}
	}
	return Meeting{}
}

func (m *Meeting) HasUser(username string) bool {
	if m.Host == username {
		return true
	} else {
		for _, member := range m.Members {
			if member.Username == username {
				return true
			}
		}
	}
	return false
}

func PrintOneMeeting(m Meeting) {
	fmt.Printf("%v\t%v\t%v\t%v\t%v\n", m.Title, m.Starttime, m.Endtime,
		m.Host, m.Members)
}
