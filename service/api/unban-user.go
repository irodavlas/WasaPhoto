package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unBanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		encodeResponse(w, Msg401, http.StatusUnauthorized)
		return
	}
	targetUsername := ps.ByName("username")
	targetUser, err := rt.isUserRegistered("", targetUsername)
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	err = rt.db.RevokeBan(*targetUser, user.Id)
	if err != nil {

		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	encodeResponse(w, Msg200, http.StatusOK)
}
