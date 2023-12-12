package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var message string

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}

	params, err := decodeParams(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}

	token, err := checkUsername(params.Username)
	if err != nil {
		var id string = generateId()

		user := User{
			Username: params.Username,
			Id:       id,
		}
		Users = append(Users, user)
		Profiles[id] = UserProfile{
			User:      user,
			Post:      make([]Photo, 0),
			Follower:  make([]User, 0),
			Following: make([]User, 0),
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(id)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
	return
}
