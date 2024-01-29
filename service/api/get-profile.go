package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//no query paramters just the token
func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")

	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		message := "Session token not valid"
		encodeResponse(w, message, http.StatusUnauthorized)
		return
	}
	err = r.ParseForm()
	if err != nil {
		message := "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}
	username, err := decodeQueryParamsUsername(r)
	if err != nil {
		message := "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}
	targetUser, err := rt.isUserRegistered("", username)
	if err != nil {
		message := "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}
	targetUser, err = rt.db.GetProfile(*targetUser, user.Id)
	if err != nil {
		message := "internal server error"
		encodeResponse(w, message, http.StatusInternalServerError)
		return
	}
	encodeResponse(w, targetUser, http.StatusOK)

}
