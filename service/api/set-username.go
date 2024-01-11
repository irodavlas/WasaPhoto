package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setNewUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message string

	id := ps.ByName("userID")

	_, err := checkId(id)
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
	p, err := decodeParams(r)
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err = encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}

	changeUsername(id, p.Username) //doesnt change the username in the Profiles
	err = encodeResponse(w, p.Username, http.StatusOK)
	if err != nil {
		panic(err)
	}
}
