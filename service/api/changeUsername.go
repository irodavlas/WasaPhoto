package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	_ = r.ParseForm()
	username := r.FormValue("username")
	token := r.Header.Get("Authorization")
	userId := ps.ByName("userId")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	_, err := rt.db.RetrieveId(username)
	if !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusConflict)
		ctx.Logger.Info("409, username " + username + " already chosen")
		return
	}

	err = rt.db.SetName(userId, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}
	ctx.Logger.Info(w, "username changed into "+username)
	_ = json.NewEncoder(w).Encode(username)

}
