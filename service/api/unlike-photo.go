package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unLikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var message string

	id := ps.ByName("userID")
	user, err := checkId(id)

	if err != nil {
		message = "No user found for those credentials"
		err := encodeResponse(w, message, http.StatusUnauthorized)
		if err != nil {
			panic(err)
		}
		return
	}

	photoId := ps.ByName("photoID")
	if err = Profiles[user.Id].checkPost(photoId); err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}

	username := r.FormValue("username")
	liker, err := isUserRegistered(username)
	if err != nil {
		message = "The server cannot or will not process the request due to an apparent client error"
		err := encodeResponse(w, message, http.StatusBadRequest)
		if err != nil {
			panic(err)
		}
		return
	}
	err = Profiles[user.Id].checkLikes(photoId, liker)
	if err != nil {
		/*
			i := findIndex(Profiles[user.Id].Post[photoId].Likes, liker.Id)
			Profiles[user.Id].Post[photoId].Likes = removeItem(Profiles[user.Id].Post[photoId].Likes, i)
		*/
	}

	//with the photo id search in the map and get the actual pointer to the photo
	message = "Success"
	err = encodeResponse(w, message, http.StatusOK)
	if err != nil {
		panic(err)
	}
}
