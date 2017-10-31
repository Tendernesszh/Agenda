package cmd

import(
  "io/ioutil"
  "testing"
  "encoding/json"
	"os"
  "github.com/Agenda/entity"
)
var curUser entity.User
func Testdestroy(t *testing.T) {
  name, _ := ioutil.ReadFile("curUser.txt")
  if name == nil {
    t.Errorf("fail")
  }
  file, err := os.OpenFile("user.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Errorf("fail")
	}
  jsonDecoder := json.NewDecoder(file)
	for jsonDecoder.More() {

		err := jsonDecoder.Decode(&curUser)
		if err != nil {
			t.Errorf("fail")
		}
	}

  destroyCmd.Run(destroyCmd, make([]string, 0))

  users := entity.GetUsers()

	for _, user := range users {
		if user.Username == curUser.Username {
			t.Errorf("fail")
		}
	}

}
