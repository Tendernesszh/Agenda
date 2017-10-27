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
	file, err := os.OpenFile(UserPath, os.O_RDONLY, 0755)
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

func AddOneUser(u User) {
	// TODO: Add one user to the json file.
}
