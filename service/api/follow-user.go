package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message string
	targetUser := new(User)
	user := new(User)
	id := ps.ByName("userID")
	_, err := checkId(id)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		message = "Authorization has been refused for those credentials"
		json.NewEncoder(w).Encode(message)
		return
	}

	targetUser.Username = r.FormValue("username")
	targetUser.Id, err = checkUsername(targetUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}
	//targetUser already in following list
	if err = checkFollowing(id, targetUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}
	userProfile := Profiles[id]
	userProfile.Following = append(userProfile.Following, *targetUser)
	Profiles[id] = userProfile

	user.Id = id // from target user add to followers list old user
	targetProfile := Profiles[targetUser.Id]
	targetProfile.Follower = append(targetProfile.Follower)
	w.WriteHeader(http.StatusOK)
	message = "Success"
	json.NewEncoder(w).Encode(message)

}
