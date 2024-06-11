package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unLikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")

	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		encodeResponse(w, Msg401, http.StatusUnauthorized)
		return
	}

	postId := ps.ByName("postId")

	err = rt.db.RemoveLike(postId, user.Id)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}

	encodeResponse(w, Msg200, http.StatusOK)

}
