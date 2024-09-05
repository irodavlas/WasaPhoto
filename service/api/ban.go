package api

import (
	"database/sql"
	"encoding/json"
	"errors"

	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	username := r.URL.Query().Get("username")
	token := r.Header.Get("Authorization")

	userId := ps.ByName("userId")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	toBanUserId, err := rt.db.RetrieveId(username)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Info("404, username " + username + notExist)
		return
	}

	err = rt.db.CheckBan(toBanUserId, userId)
	if !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusConflict)
		ctx.Logger.Info("409, username " + username + " already banned")
		return
	}

	err = rt.db.BanUser(userId, toBanUserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}
	ctx.Logger.Info(w, "201, user "+username+" banned")
	_ = json.NewEncoder(w).Encode(toBanUserId)
}
