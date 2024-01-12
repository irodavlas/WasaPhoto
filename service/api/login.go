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

	user, err := decodeParams(r)
	if err != nil {
		fmt.Println(err)
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}

	token, err := checkUsername(user.Username)
	if err != nil {
		var id string = generateId()

		user.Id = id

		userProfile := UserProfile{
			User: &User{
				Username: user.Username,
				Id:       user.Id,
			},
			Post:      make([]Photo, 0),
			Follower:  make(map[string]*User),
			Following: make(map[string]*User),
		}

		Profiles[user.Id] = &userProfile

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
