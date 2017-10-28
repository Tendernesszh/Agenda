package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type SimpleUser struct {
	Username string `json:"username"`
}

var (
	UserPath string = "data/users.json"
	userList []User
)

func init() {
	file, err := os.OpenFile(UserPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonDecoder := json.NewDecoder(file)
	for jsonDecoder.More() {
		var curUser User
		err := jsonDecoder.Decode(&curUser)
		if err != nil {
			fmt.Println(err)
			return
		}
		userList = append(userList, curUser)
	}
}

func GetUsers() []User {
	return userList
}

func DeleteOneUser(username string) { //input the username of the user to be deleted
	Userfile, err := os.OpenFile(UserPath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	AllMeetings := GetMeetings()
	for meetingIndex, meeting := range AllMeetings {
		for i, user := range meeting.Members {
			if user.Username == username {
				meeting.Members = append(meeting.Members[:i], meeting.Members[i+1:]...)
				if len(meeting.Members) == 0 {
					AllMeetings = append(AllMeetings[:meetingIndex], AllMeetings[meetingIndex+1:]...)
				}
			}
		}
	}

	AllUser := GetUsers()
	Meetingfile, err := os.OpenFile(MeetingPath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, user := range AllUser {
		if user.Username == username {
			AllUser = append(AllUser[:i], AllUser[i+1:]...)
			break
		}
	}
	jsonEncoder := json.NewEncoder(Userfile)
	for _, u := range AllUser {
		encodeErr := jsonEncoder.Encode(&u)
		if encodeErr != nil {
			fmt.Println(encodeErr)
			return
		}
	}
	jsonEncoder = json.NewEncoder(Meetingfile)
	for _, m := range AllMeetings {
		encodeErr := jsonEncoder.Encode(&m)
		if encodeErr != nil {
			fmt.Println(encodeErr)
			return
		}
	}
	// TODO: call logout

}
func AddOneUser(u User) {
	// Add one user to the json file.
	file, err := os.OpenFile(UserPath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonEncoder := json.NewEncoder(file)
	encodeErr := jsonEncoder.Encode(&u)
	if encodeErr != nil {
		fmt.Println(encodeErr)
		return
	}
}
