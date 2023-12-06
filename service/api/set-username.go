package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setNewUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message string

	id := ps.ByName("userID")
	_, err := checkId(id)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}

	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}
	p, err := decodeParams(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = "The server cannot or will not process the request due to an apparent client error"
		json.NewEncoder(w).Encode(message)
		return
	}

	changeUsername(id, p.Username)
	w.WriteHeader(http.StatusOK)
	message = "Success"
	json.NewEncoder(w).Encode(message)

}
