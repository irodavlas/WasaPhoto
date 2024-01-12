package api

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

var Profiles = make(map[string]*UserProfile)

func checkLenght(username string) bool {
	if len(username) < 3 || len(username) > 16 {
		return true
	}
	return false
}
func decodeParams(r *http.Request) (*User, error) {
	user := new(User)
	username := r.FormValue("username")
	if checkLenght(username) {
		return nil, errors.New("invalid params")
	}
	user.Username = username
	return user, nil
}
func checkId(id string) (*User, error) {
	if Profiles[id] != nil {
		return Profiles[id].User, nil
	}
	return nil, errors.New("username not found")
}
func checkUsername(username string) (*User, error) {
	for _, profile := range Profiles {
		if profile.User.Username == username {

			return profile.User, nil
		}
	}
	return nil, errors.New("username not found")
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

func uploadPhotoParams(r *http.Request, id string) *Photo {

	pic := new(Photo) //new returns a pointer to struct Photo
	pic.PhotoID = len(Profiles[id].Post)
	pic.Img = ""
	pic.OwnerID = id
	pic.UpDate = time.Now().GoString()

	return pic
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

func (P *User) changeUsername(username string) {
	P.Username = username
}
func (Profile *UserProfile) checkFollowing(target *User) error {
	if Profile.Following[target.Id] != nil {
		return errors.New("already following")
	}
	return nil
}
