package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := r.ParseForm()
	if err != nil {

		message := "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		fmt.Println(err)
		return
	}

	username, err := decodeQueryParams(r)
	if err != nil {
		message := "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		fmt.Println(err)
		return
	}

	token, err := isUserRegistered(username)
	if err != nil {

		id := generateGenericToken()

		//create connection with database here to store each collumn

		user := NewUser(id, username)

		err := encodeResponse(w, user, http.StatusCreated)
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
