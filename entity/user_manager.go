package entity

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

const USER_PATH string = "data/users.json"

func GetUsers() []User {
	userList := make([]User, 0)
	file, err := os.OpenFile(USER_PATH, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	jsonDecoder := json.NewDecoder(file)
	for jsonDecoder.More() {
		var curUser User
		err := jsonDecoder.Decode(&curUser)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		userList = append(userList, curUser)
	}
	return userList
}

func DeleteOneUser(username string) { //input the username of the user to be deleted
	Userfile, err := os.OpenFile(USER_PATH, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	AllMeetings := GetMeetings()
	newMeetings := make([]Meeting, 0)
	for _, meeting := range AllMeetings {
		del := false
		if meeting.Host == username {
			del = true
		}
		newMembers := make([]SimpleUser, 0)
		delMember := false
		for _, user := range meeting.Members {
			if user.Username == username {
				delMember = true
				if len(meeting.Members) == 1 {
					del = true
				}
				break
			}
			if !delMember {
				newMembers = append(newMembers, user)
			}
		}
		if !del {
			newMeetings = append(newMeetings, meeting)
		}
	}
	UpdateMeeting(newMeetings)

	AllUser := GetUsers()

	for i, user := range AllUser {
		if user.Username == username {
			AllUser = append(AllUser[:i], AllUser[i+1:]...)
			break
		}
	}
	os.Truncate(USER_PATH, 0)
	jsonEncoder := json.NewEncoder(Userfile)
	for _, u := range AllUser {
		encodeErr := jsonEncoder.Encode(&u)
		if encodeErr != nil {
			fmt.Println(encodeErr)
			return
		}
	}
}
func AddOneUser(u User) {
	// Add one user to the json file.
	file, err := os.OpenFile(USER_PATH, os.O_WRONLY|os.O_APPEND, os.ModePerm)
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
