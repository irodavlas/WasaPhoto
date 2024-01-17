package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unBanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var message string
	targetUser := new(User)
	user := new(User)

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
	if err = Profiles[user.Id].checkBanList(targetUser); err == nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}
	delete(Profiles[user.Id].Banned, targetUser.Id)
	message = "Success"
	err = encodeResponse(w, message, http.StatusOK)
	if err != nil {
		panic(err)
	}
}
