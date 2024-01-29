package api

import (
	"myproject/service/types"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//sauthorizationHeader := r.Header.Get("Authorization")

	err := r.ParseForm()
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
	user, err := rt.isUserRegistered("", username)
	if err != nil {
		user := types.NewUser(generateGenericToken(), username)

		//db function call
		err = rt.db.InsertUser(*user)
		if err != nil {
			message := "Internal server error"
			encodeResponse(w, message, http.StatusInternalServerError)
			return
		}
		encodeResponse(w, user, http.StatusCreated)
		return
	}
	encodeResponse(w, user.Id, http.StatusOK)
}
