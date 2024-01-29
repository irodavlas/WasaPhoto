package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	postId := decodeQueryParamsPostId(r)
	err = rt.db.RemovePost(user.Id, postId)
	if err != nil {
		message := "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}

	imagesDir := "./service/images"
	fileToDelete := postId + ".jpg"
	filePathToDelete := filepath.Join(imagesDir, fileToDelete)
	err = os.Remove(filePathToDelete)
	if err != nil {
		fmt.Println("Error:", err)
	}
	message := "Success"
	encodeResponse(w, message, http.StatusOK)

}
