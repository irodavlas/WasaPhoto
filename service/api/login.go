package api

import (
	"myproject/service/types"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	username, err := decodeQueryParamsUsername(r)
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	user, err := rt.isUserRegistered("", username)
	if err != nil {
		user := types.NewUser(generateGenericToken(), username)

		err = rt.db.InsertUser(*user)
		if err != nil {

			encodeResponse(w, Msg500, http.StatusInternalServerError)
			return
		}
		encodeResponse(w, user.Id, http.StatusCreated)
		return
	}
	encodeResponse(w, user.Id, http.StatusOK)
}
