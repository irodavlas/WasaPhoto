package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message string
	username := r.FormValue("username")

	id, err := checkUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}
	jsonData, err := json.Marshal(Profiles[id])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message = "Internal Server Error"
		json.NewEncoder(w).Encode(message)
	}

	w.WriteHeader(http.StatusOK)
	message = "Success: " + string(jsonData)
	json.NewEncoder(w).Encode(message)

}
