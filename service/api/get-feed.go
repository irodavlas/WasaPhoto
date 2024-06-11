package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// no query paramters just the token
func (rt *_router) getFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")

	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		encodeResponse(w, Msg401, http.StatusUnauthorized)
		return
	}

	feed, err := rt.db.GetFeed(*user)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	feed.Feed = sortPosts(*feed)
	encodeResponse(w, feed.Feed, 200)
}
