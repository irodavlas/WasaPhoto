package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//no query paramters just the token
func (rt *_router) getFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")

	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		message := "Session token not valid"
		encodeResponse(w, message, http.StatusUnauthorized)
		return
	}

	feed, err := rt.db.GetFeed(*user)
	if err != nil {
		message := "Internal server error"
		encodeResponse(w, message, http.StatusInternalServerError)
		return
	}
	feed.Feed = sortPosts(*feed)
	encodeResponse(w, feed.Feed, 200)
}
