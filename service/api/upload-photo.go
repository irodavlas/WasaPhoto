package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var message string
	id := ps.ByName("userID")
	_, err := checkId(id)
	if err != nil {
		message = "Authorization has been refused for those credentials"
		err := encodeResponse(w, message, http.StatusUnauthorized)
		if err != nil {
			panic(err)
		}
		return
	}
	err = r.ParseForm()
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}
	pic := uploadPhotoParams(r, id)
	userProfile := Profiles[id]
	userProfile.Post = append(userProfile.Post, *pic)
	//implement more logic for the images
	Profiles[id] = userProfile
	message = "photo uploaded succesfully"
	err = encodeResponse(w, message, http.StatusOK)
	if err != nil {
		panic(err)
	}
}
