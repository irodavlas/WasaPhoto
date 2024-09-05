package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	userId := r.URL.Query().Get("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	photoId := r.URL.Query().Get("photoId")

	err := rt.db.CheckPhotoExistence(photoId)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Info("404, photo " + photoId + notExist)
		return
	}

	err = rt.db.CheckLikeExistenceFromPhoto(photoId, userId)
	if !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusConflict)
		ctx.Logger.Info("409, photo already liked")
		return
	}

	likeId := generateId()
	for !errors.Is(err, sql.ErrNoRows) {
		likeId = generateId()
		err = rt.db.CheckLikeExistence(likeId)
	}

	err = rt.db.PutLike(likeId, userId, photoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	ctx.Logger.Info("201, like succesfully put")
	_ = json.NewEncoder(w).Encode(likeId)

}
