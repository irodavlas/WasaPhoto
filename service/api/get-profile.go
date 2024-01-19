package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var message string
	username := r.FormValue("username")
	user, err := rt.isUserRegistered("", username)
	print(user)
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err = encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}
	//redefine response
	err = encodeResponse(w, user, http.StatusOK)
	if err != nil {
		message = "internal server error"
		encodeResponse(w, message, http.StatusInternalServerError)
		if err != nil {
			panic(err)
		}
	}

}
