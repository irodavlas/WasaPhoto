package api

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

var Users = make([]User, 0)
var Profiles = make(map[string]UserProfile)

func checkLenght(username string) bool {
	if len(username) < 3 || len(username) > 16 {
		return true
	}
	return false
}
func decodeParams(r *http.Request) (*Params, error) {
	p := new(Params)
	username := r.FormValue("username")
	if checkLenght(username) {
		return nil, errors.New("invalid params")
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
	for _, elem := range Profiles {
		if elem.User.Id == id {
			//cant assign to struct profiles

		}
	}
}

func uploadPhotoParams(r *http.Request, id string) *Photo {

	pic := new(Photo) //new returns a pointer to struct Photo
	pic.PhotoID = len(Profiles[id].Post)
	pic.Img = ""
	pic.OwnerID = id
	pic.UpDate = time.Now().GoString()

	return pic
}
func checkFollowing(userId string, target *User) error {
	for _, x := range Profiles[userId].Following {
		if x.Username == target.Username {
			return errors.New("already following this user")
		}
	}
	return nil
}
func encodeResponse[T any](w http.ResponseWriter, message T, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data := map[string]interface{}{
		"code":    statusCode,
		"message": message,
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	jsonData = append(jsonData, '\n') //just for sake of readability
	w.Write(jsonData)                 //Grants better costumization over encoding, the input is []byte
	return nil

}
