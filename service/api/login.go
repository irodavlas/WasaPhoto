package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var message string

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}

	params, err := decodeParams(r)
	if err != nil {
		fmt.Println(err)
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
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
		err := encodeResponse(w, id, http.StatusCreated)
		if err != nil {
			panic(err)
		}
		return
	}

	err = encodeResponse(w, token, http.StatusOK)
	if err != nil {
		panic(err)
	}

}
