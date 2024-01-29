package api

import (
	"fmt"
	"io/ioutil"
	"myproject/service/types"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	authorizationHeader := r.Header.Get("Authorization")
	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		message := "Session token not valid"
		encodeResponse(w, message, http.StatusUnauthorized)
		return
	}
	//image is sent within the header
	//later saved in a dir, named as its Post Id
	// Determine the file extension based on the Content-Type header
	PhotoId := generateGenericToken()
	//downloadDirectory := "/service/images" // relative path from service/api/filethatsavesit
	downloadDirectory := "./service/images"

	path := filepath.Join(downloadDirectory, PhotoId+".jpg")
	_, err = os.Create(path)
	if err != nil {
		encodeResponse(w, err, http.StatusInternalServerError)
		fmt.Println("Error creating file:", err)
		return
	}

	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		message := "The server cannot or will not process the request due to an apparent client error"
		encodeResponse(w, message, http.StatusBadRequest)
		return
	}
	err = saveResponseBodyToFile(responseBody, path)
	if err != nil {
		fmt.Println("Error saving response body to file:", err)
		encodeResponse(w, err, http.StatusInternalServerError)
		return
	}
	post := types.NewPost(PhotoId, user.Id, path, user.Username)
	post.OwnerUsername = user.Username
	err = rt.db.InsertPost(*post)
	if err != nil {
		fmt.Println("Error saving post information into db:", err)
		encodeResponse(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	message := "post uploaded succesfully"
	encodeResponse(w, message, http.StatusCreated)
}
