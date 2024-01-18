package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

var Profiles = make(map[string]User)

func checkLenght(username string) bool {
	if len(username) < 3 || len(username) > 16 {
		return true
	}
	return false
}
func decodeQueryParams(r *http.Request) (string, error) {

	parameter := r.FormValue("username")
	if checkLenght(parameter) {
		return "", errors.New("invalid params")
	}

	return parameter, nil
}
func checkId(id string) (*User, error) {
	/*
		if Profiles[id] != nil {
			return Profiles[id].User, nil
		}
		return nil, errors.New("username not found")
	*/
	return nil, nil
}
func (rt *_router) isUserRegistered(username string) (*User, error) {
	//retrive info from db
	//make return type to be a Use, for working use now ill write string

	user, err := rt.db.GetUser(username)
	if err != nil {
		return nil, err
	}
	if username == user.Username {
		return &user, nil
	}

	return nil, errors.New("username not found")
}

func generateGenericToken() string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	var s string
	for i := 0; i < 10; i++ {
		index := rand.Intn(len(letters))
		s += string(letters[index])
	}
	return s

}

func uploadPhotoParams(id string) *Post {
	/*
		pic := new(Photo) //new returns a pointer to struct Photo
		pic.PhotoID = generateGenericToken()
		pic.Img = ""
		pic.OwnerID = id
		pic.UpDate = time.Now().GoString()
		pic.Likes = make([]*User, 0)
		pic.Comments = make([]*Comment, 0)
		return pic
	*/
	return new(Post)
}

func removeItem[T any](slice []T, index int) []T {
	// Check if the index is valid
	if index < 0 || index >= len(slice) {
		fmt.Println("Invalid index")
		return slice
	}

	// Remove the element at the specified index
	slice = append(slice[:index], slice[index+1:]...)

	return slice
}
func findIndex(slice []*User, id string) int {
	for i, user := range slice {
		if user.Id == id {
			return i
		}
	}
	return -1
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
func (Profile User) checkFollowing(target *User) error {
	/*
		if Profile.Following[target.Id] != nil || Profile.User.Id == target.Id {
			return errors.New("already following")
		}
		return nil
	*/
	return nil
}

func (Profile User) checkBanList(target *User) error {
	/*
		if Profile.Banned[target.Id] != nil || Profile.User.Id == target.Id {
			return errors.New("already in banned list")
		}
		return nil
	*/
	return nil
}

func (Profile User) checkPost(picId string) error {
	/*
		if Profile.Post[picId] == nil {
			return errors.New("Post not found")
		}
		return nil
	*/
	return nil
}
func (Profile User) checkLikes(picId string, user *User) error {
	/*
		for _, el := range Profile.Post[picId].Likes {
			if el.Id == user.Id {
				return errors.New("post already liked")
			}
		}
		return nil
	*/
	return nil
}
