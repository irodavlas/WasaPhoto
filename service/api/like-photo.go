package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	err = rt.db.InsertLike(generateGenericToken(), postId, *user)
	if err != nil {
		message := "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}
	message := "Success"
	encodeResponse(w, message, http.StatusOK)
}
