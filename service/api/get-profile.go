package api

import (
	"myproject/service/types"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")

	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		encodeResponse(w, Msg401, http.StatusUnauthorized)
		return
	}
	/*
		when no username is provided then the user profile is fetched, if a username is fetched the first user to match
		will be displayed
	*/
	username := r.URL.Query().Get("username")
	var targetUser *types.User
	if username != "" {
		targetUser, err = rt.isUserRegistered("", username)
		if err != nil {
			rt.baseLogger.Println(err)
			return
		}
	} else {
		targetId := ps.ByName("userId")
		targetUser, err = rt.isUserRegistered(targetId, "")
	}

	if err != nil {
		rt.baseLogger.Println(err)
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	targetUser, err = rt.db.GetProfile(*targetUser, user.Id)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	encodeResponse(w, targetUser, http.StatusOK)

}
