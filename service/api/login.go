package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func generateId() string {
	length := 16
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	output := ""

	for i := 0; i < length; i++ {
		random_numb := rand.Intn(100) % len(charset)
		output += string(charset[random_numb])
	}

	return output
}

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	_ = r.ParseForm()
	username := r.FormValue("username")
	userId, err := rt.db.RetrieveId(username)

	// register a new user
	if err != nil {
		ctx.Logger.Info("user " + username + " not registered")

		if len(username) < 3 || len(username) > 16 {
			w.WriteHeader(http.StatusBadRequest)
			ctx.Logger.Info("400, bad request")
			return
		}

		userId := generateId()
		err = rt.db.CheckUserExistence(userId)
		for !errors.Is(err, sql.ErrNoRows) {
			userId = generateId()
			err = rt.db.CheckUserExistence(userId)
			ctx.Logger.Info("login ", err)
		}

		err := rt.db.RegisterUser(username, userId)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			ctx.Logger.Info("500, internal server error")
			return
		}

		ctx.Logger.Info("user " + username + " is now correctly registered")
		_ = json.NewEncoder(w).Encode(userId)

	} else {
		ctx.Logger.Info("user " + username + " logged in")
		_ = json.NewEncoder(w).Encode(userId)
	}
}

func (rt *_router) authorize(userId string, token string) bool {
	if len(token) < 6 {
		return false
	}
	token = token[7:]

	return rt.db.CheckUserExistence(token) == nil && userId == token
}
