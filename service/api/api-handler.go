package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	rt.router.POST("/users/:userID", rt.setNewUsername)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.POST("/doLogin", rt.doLogin)
	rt.router.POST("/users/:userID/photo", rt.uploadPhoto)
	rt.router.GET("/users/profile", rt.getProfile)
	rt.router.POST("/users/:userID/followers", rt.followUser)
	rt.router.DELETE("/users/:userID/followers", rt.unfollowUser)
	rt.router.POST("/users/:userID/ban", rt.banUser)
	rt.router.DELETE("/users/:userID/ban", rt.unBanUser)
	rt.router.POST("/photos/:userID/:photoID/likes", rt.likePhoto)
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
