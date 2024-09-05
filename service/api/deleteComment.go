package api

import (
	"database/sql"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	userId := r.URL.Query().Get("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, anauthorized")
		return
	}

	commentId := ps.ByName("commentId")
	err := rt.db.CheckCommentExistence(commentId)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Info("404, comment: " + commentId + notExist)
		return
	}

	err = rt.db.CheckCommentOwnership(commentId, userId)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Info("403, you are not allowed to delete this comment")
		return
	}

	err = rt.db.DeleteComment(commentId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info(w, "500, internal server error")
	}
	ctx.Logger.Info("200, comment: " + commentId + " deleted")
}
