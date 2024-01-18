package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setNewUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var message string
	user := new(User)
	id := ps.ByName("userID")

	user, err := checkId(id)
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err = encodeResponse(w, message, http.StatusUnauthorized)
		if err != nil {
			panic(err)
		}
		return
	}

	err = r.ParseForm()
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err = encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}
	newUsername, err := decodeQueryParams(r)
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err = encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}

	user.changeUsername(newUsername)

	err = encodeResponse(w, user.Username, http.StatusOK)
	if err != nil {
		panic(err)
	}
}
