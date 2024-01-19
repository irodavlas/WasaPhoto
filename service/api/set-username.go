package api

import (
	"myproject/service/types"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setNewUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var message string
	user := new(types.User)
	user.Id = ps.ByName("userID")

	err := r.ParseForm()
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err = encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}
	user.Username, err = decodeQueryParams(r)
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err = encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}
	err = rt.db.InsertUser(*user)
	if err != nil {
		panic(err)
	}
	err = encodeResponse(w, user.Username, http.StatusOK)
	if err != nil {
		panic(err)
	}
}
