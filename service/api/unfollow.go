package api

import (
	"database/sql"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	_ = r.ParseForm()
	w.Header().Set("content-type", "application/json")
	userId := ps.ByName("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	username := ps.ByName("username")
	followingUserId, err := rt.db.RetrieveId(username)

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Info("404, user " + username + notExist)
		return
	}
	err = rt.db.CheckFollow(userId, followingUserId)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Info("404, not following " + username)
		return
	}

	err = rt.db.DeleteFollow(userId, followingUserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}
	ctx.Logger.Info("200, user " + username + " not following anymore")
}
