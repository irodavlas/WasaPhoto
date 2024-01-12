package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var message string
	user := new(User)
	targetUser := new(User)

	//user making the request
	id := ps.ByName("userID")
	user, err := checkId(id)
	if err != nil {
		message = "Authorization has been refused for those credentials"
		err := encodeResponse(w, message, http.StatusUnauthorized)
		if err != nil {
			panic(err)
		}
		return
	}

	targetUser.Username = r.FormValue("username")
	targetUser, err = checkUsername(targetUser.Username)
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}

	delete(Profiles[user.Id].Following, targetUser.Id)
	delete(Profiles[targetUser.Id].Follower, user.Id)

	message = "Success"
	err = encodeResponse(w, message, http.StatusOK)
	if err != nil {
		panic(err)
	}

}
