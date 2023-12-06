package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	for _, x := range Users {
		println(x.Id, x.Username)
	}
	json.NewEncoder(w).Encode("message")
}
