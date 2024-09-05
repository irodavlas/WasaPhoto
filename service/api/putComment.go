package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Info("400, bad request")
		return
	}

	userId := r.URL.Query().Get("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}

	photoId := r.URL.Query().Get("photoId")
	content := r.FormValue("content")
	err = rt.db.CheckPhotoExistence(photoId)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Info("404, photoId: " + photoId + notExist)
		return
	}

	commentId := generateId()
	err = rt.db.CheckCommentExistence(commentId)
	for !errors.Is(err, sql.ErrNoRows) {
		commentId = generateId()
		err = rt.db.CheckCommentExistence(commentId)
	}

	err = rt.db.PutComment(commentId, userId, photoId, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	ctx.Logger.Info("201, comment succesfully posted")
	_ = json.NewEncoder(w).Encode(commentId)

}
