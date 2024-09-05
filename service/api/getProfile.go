package api

import (
	"database/sql"
	"encoding/json"
	"errors"

	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	username := ps.ByName("username")
	userId := r.URL.Query().Get("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	toVisitUserId, err := rt.db.RetrieveId(username)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Info("404, user not found")
		return
	}

	images, err := rt.db.GetProfile(userId, toVisitUserId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Info("500, internal sever error")
		return
	}
	if len(images) == 0 {
		ctx.Logger.Info("200, profile retrieved but empty")

	} else {
		ctx.Logger.Info("200, profile succesfully retrieved")
	}

	followers, err := rt.db.RetrieveFollowers(toVisitUserId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.Info("500, internal sever error")
		return
	}

	following, err := rt.db.RetrieveFollowing(toVisitUserId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.Info("500, internal sever error")
		return
	}
	var banned bool
	err = rt.db.CheckBan(userId, toVisitUserId)
	if errors.Is(err, sql.ErrNoRows) {
		banned = false
	} else {
		banned = true
	}

	profile := Profile{followers, following, images, banned}

	_ = json.NewEncoder(w).Encode(profile)
}
