package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	userId := ps.ByName("userId")
	refresh := r.URL.Query().Get("refresh")
	token := r.Header.Get("Authorization")

	update := (refresh == "true")
	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	var fullStream []database.Photo
	if update {
		var err error
		fullStream, err = rt.db.GetFullStream(userId)
		if !errors.Is(err, sql.ErrNoRows) && err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.Info("500, internal server error")
			return
		}
	} else {
		fullStream = streamDatabase[userId]
	}

	var streamToEncode []database.Photo
	if len(fullStream) >= 15 {
		streamToEncode = fullStream[:15]
		streamDatabase[userId] = fullStream[15:]
	} else {
		streamToEncode = fullStream
		streamDatabase[userId] = make([]database.Photo, 0)
	}

	_ = json.NewEncoder(w).Encode(streamToEncode)
	ctx.Logger.Info("200, stream successfully retrieved")

}
