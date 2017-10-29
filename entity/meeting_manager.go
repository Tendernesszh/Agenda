package entity

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

const MEETING_PATH string = "data/meetings.json"

func UpdateMeeting(meetings []Meeting) {
	file, err := os.OpenFile(MEETING_PATH, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Truncate(MEETING_PATH, 0)
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
	file, err := os.OpenFile(MEETING_PATH, os.O_WRONLY|os.O_APPEND, os.ModePerm)
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
	meetingList := make([]Meeting, 0)
	file, err := os.OpenFile(MEETING_PATH, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	jsonDecoder := json.NewDecoder(file)
	for jsonDecoder.More() {
		var curMeeting Meeting
		err := jsonDecoder.Decode(&curMeeting)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		meetingList = append(meetingList, curMeeting)
	}
	return meetingList
}

func GetMeeting(title string) Meeting {
	meetingList := GetMeetings()
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
