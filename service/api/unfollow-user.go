package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message string
	user := new(User)
	id := ps.ByName("userID")
	_, err := checkId(id)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		message = "Authorization has been refused for those credentials"
		json.NewEncoder(w).Encode(message)
		return
	}

	user.Username = r.FormValue("username")
	user.Id, err = checkUsername(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}
	if err = checkFollowing(id, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}
	//FollowingList := Profiles[id]

}
