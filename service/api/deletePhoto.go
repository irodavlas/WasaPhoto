package api

import (
	"database/sql"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	userId := r.URL.Query().Get("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info(w, "401, unauthorized")
		return
	}

	photoId := ps.ByName("photoId")
	err := rt.db.CheckPhotoExistence(photoId)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Info("404, the photo" + notExist)
		return
	}

	err = rt.db.CheckPhotoOwnership(photoId, userId)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Info("403, you are not allowed to delete this photo")
		return
	}

	err = rt.db.DeleteLikeFromPhoto(photoId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}

	err = rt.db.DeleteCommentFromPhoto(photoId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}

	err = rt.db.DeletePhoto(photoId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}

	ctx.Logger.Info("200, photo " + photoId + " deleted")
}
