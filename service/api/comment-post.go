package api

import (
	"io/ioutil"
	"myproject/service/types"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	//commment content is passed as raw text in the body request
	var comment types.Comment
	comment.Username = user.Username
	comment.Message = string(body)

	err = rt.db.InsertComment(generateGenericToken(), postId, user.Id, comment.Message)
	if err != nil {
		message := "Internal server error"
		encodeResponse(w, message, http.StatusInternalServerError)
		return
	}
	message := "Success"
	encodeResponse(w, message, http.StatusOK)

}
