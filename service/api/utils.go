package api

import (
	"errors"
	"math/rand"
	"net/http"
)

var Users = make([]User, 0)

func checkLenght(username string) bool {
	if len(username) < 3 || len(username) > 16 {
		return true
	}
	return false
}
func decodeParams(r *http.Request) (*Params, error) {
	p := new(Params)
	username := r.FormValue("name")
	if checkLenght(username) {
		return nil, errors.New("Invalid params")
	}
	p.Username = username
	return p, nil
}
func checkId(id string) (string, error) {
	for _, x := range Users {
		if x.Id == id {
			return x.Id, nil
		}
	}
	return "", errors.New("username not found")
}
func checkUsername(username string) (string, error) {
	for _, x := range Users {
		if x.Username == username {
			return x.Id, nil
		}
	}
	return "", errors.New("username not found")
}

func generateId() string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	var s string
	for i := 0; i < 10; i++ {
		index := rand.Intn(len(letters))
		s += string(letters[index])
	}
	return s

}
func changeUsername(id string, new string) {
	for i, elem := range Users {
		if elem.Id == id {
			Users[i].Username = new
			print(Users[i].Username)
		}
	}
}
