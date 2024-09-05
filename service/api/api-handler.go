package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Predefined Routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	// Login
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	// Settings
	rt.router.PUT("/users/:userId", rt.wrap(rt.setMyUsername))
	// View
	rt.router.GET("/profiles/:username", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:userId/stream", rt.wrap(rt.getMyStream))
	// Network
	rt.router.POST("/users/:userId/followed/", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userId/followed/:username", rt.wrap(rt.unfollowUser))
	// Photos
	rt.router.POST("/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:photoId", rt.wrap(rt.deletePhoto))
	// rt.router.GET("/photos/:photoId", rt.wrap(rt.getImage))
	// Bans
	rt.router.POST("/users/:userId/banned/", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userId/banned/:username", rt.wrap(rt.unbanUser))
	// Likes
	rt.router.POST("/likes/", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/likes/:likeId", rt.wrap(rt.unlikePhoto))
	// Comments
	rt.router.POST("/comments/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/comments/:commentId", rt.wrap(rt.uncommentPhoto))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
