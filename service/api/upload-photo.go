package api

import (
	"myproject/service/types"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var message string
	user := new(types.User)

	id := ps.ByName("userID")

	err := r.ParseForm()
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}
	user, err = rt.isUserRegistered(id, "")
	if err != nil {
		panic(err)
	}
	pic := uploadPhotoParams(user.Id)

	print(pic)
	/*
		Profiles[user.Id].Post[pic.PhotoID] = pic
	*/
	message = "photo uploaded succesfully"
	err = encodeResponse(w, message, http.StatusOK)
	if err != nil {
		panic(err)
	}
}
