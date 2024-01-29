package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"myproject/service/types"
	"sort"
	"strings"

	"net/http"
)

var Profiles = make(map[string]types.User)

func checkLenght(username string) bool {
	if len(username) < 3 || len(username) > 16 {
		return true
	}
	return false
}
func decodeQueryParamsUsername(r *http.Request) (string, error) {

	parameter := r.FormValue("username")
	if checkLenght(parameter) {
		return "", errors.New("invalid params")
	}
	return parameter, nil
}
func decodeQueryParamsCommentId(r *http.Request) string {
	parameter := r.FormValue("commentId")
	return parameter
}
func decodeQueryParamsPostId(r *http.Request) string {
	parameter := r.FormValue("postId")
	return parameter
}
func (rt *_router) isTokenValid(token string) (*types.User, error) {
	// Check if the Authorization header starts with "Bearer"
	if !strings.HasPrefix(token, "Bearer ") {
		return nil, errors.New("invalid Authorization header format")
	}

	// Extract the token from the Authorization header
	token = strings.TrimPrefix(token, "Bearer ")
	user, err := rt.db.GetUser(token, "")
	if err != nil {
		return nil, errors.New("Unauthorized")
	}
	return &user, nil
}

func (rt *_router) isUserRegistered(id string, username string) (*types.User, error) {
	//retrive info from db
	//make return type to be a Use, for working use now ill write string

	user, err := rt.db.GetUser(id, username)
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

func saveResponseBodyToFile(responseBody []byte, filePath string) error {
	return ioutil.WriteFile(filePath, responseBody, 0644)
}
func encodeResponse(w http.ResponseWriter, message interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data := map[string]interface{}{
		"code":    statusCode,
		"message": message,
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	jsonData = append(jsonData, '\n') //just for sake of readability
	_, _ = w.Write(jsonData)          //Grants better costumization over encoding, the input is []byte

}

func sortPosts(feed types.Feed) []types.Post {
	sort.Slice(feed.Feed, func(i, j int) bool {
		return feed.Feed[i].Time > feed.Feed[j].Time
	})

	// Print the sorted posts
	return feed.Feed
}
