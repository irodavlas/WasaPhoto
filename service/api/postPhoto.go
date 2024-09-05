package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userId := r.URL.Query().Get("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	image, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Info("400, bad request")
		return
	}

	photoId := generateId()
	err = rt.db.CheckPhotoExistence(photoId)
	for !errors.Is(err, sql.ErrNoRows) {
		photoId = generateId()
		err = rt.db.CheckPhotoExistence(photoId)
	}

	err = rt.db.PostPhoto(photoId, userId, image, time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
	}
	w.WriteHeader(http.StatusCreated)
	ctx.Logger.Info("201, post created")
	_ = json.NewEncoder(w).Encode(photoId)

}
