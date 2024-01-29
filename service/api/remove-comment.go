package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")

	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		message := "Session token not valid"
		encodeResponse(w, message, http.StatusUnauthorized)
		return
	}
	err = r.ParseForm()
	if err != nil {
		message := "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}
	//this should be passed by the fe (being the name of the pic)
	//so no error handling
	postId := decodeQueryParamsPostId(r)
	commentId := decodeQueryParamsCommentId(r)
	//gotta check wheter he owner is trying to remove the comment
	//i can do in the db by removing where the auth token is
	err = rt.db.RemoveComment(postId, user.Id, commentId)
	if err != nil {
		message := "internal server error"
		encodeResponse(w, message, http.StatusInternalServerError)
		return
	}
	message := "Success"
	encodeResponse(w, message, http.StatusOK)
}
