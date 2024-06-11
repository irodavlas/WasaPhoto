package api

import (
	"io"
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
		encodeResponse(w, Msg401, http.StatusUnauthorized)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {

		encodeResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	if fileHeader.Header.Get("Content-Type") != "image/jpeg" {
		encodeResponse(w, "file extensions allowed are : jpeg", http.StatusBadRequest)
		return
	}
	PhotoId := generateGenericToken()
	// downloadDirectory := "/service/images" // relative path from service/api/filethatsavesit
	downloadDirectory := "./webui/public/images"

	path := filepath.Join(downloadDirectory, PhotoId+".jpg")
	newFile, err := os.Create(path)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, file)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	post := types.NewPost(PhotoId, user.Id, path, user.Username)
	post.OwnerUsername = user.Username
	err = rt.db.InsertPost(*post)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}

	encodeResponse(w, post, http.StatusOK)
}
