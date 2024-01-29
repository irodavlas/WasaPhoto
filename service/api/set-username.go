package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setNewUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	var message string

	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		message = "Session token not valid"
		encodeResponse(w, message, http.StatusUnauthorized)
		return
	}

	err = r.ParseForm()
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}
	user.Username, err = decodeQueryParamsUsername(r)
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}
	//allows user to have the same name, id will alwsays be different
	err = rt.db.UpdateUsername(user.Id, user.Username)
	if err != nil {
		message := "Internal server error"
		encodeResponse(w, message, http.StatusInternalServerError)
		return
	}
	encodeResponse(w, user.Username, http.StatusOK)
}
