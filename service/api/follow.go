package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	userId := ps.ByName("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	username := r.URL.Query().Get("username")

	toFollowUserId, err := rt.db.RetrieveId(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Info("404, user " + username + " not found")
		return
	}
	err = rt.db.CheckFollow(userId, toFollowUserId)

	if !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusConflict)
		ctx.Logger.Info("409, already following")
		return
	}

	err = rt.db.FollowUser(userId, toFollowUserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}

	ctx.Logger.Info("201, now following " + username)
	_ = json.NewEncoder(w).Encode(toFollowUserId)
}
