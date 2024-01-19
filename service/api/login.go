package api

import (
	"fmt"
	"myproject/service/types"

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
	user, err := rt.isUserRegistered("", username)
	if err != nil {

		user := types.User{
			Id:       generateGenericToken(),
			Username: username,
		}
		//db function call
		err = rt.db.InsertUser(user)
		if err != nil {
			panic(err)
		}
		err := encodeResponse(w, user, http.StatusCreated)
		if err != nil {
			panic(err)
		}
		return
	}

	err = encodeResponse(w, user.Id, http.StatusOK)
	if err != nil {
		panic(err)
	}

}
