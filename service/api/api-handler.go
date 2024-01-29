package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	rt.router.POST("/users/profile/username", rt.setNewUsername)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.POST("/session", rt.doLogin)
	rt.router.GET("/users/profile", rt.getProfile)
	rt.router.POST("/users/profile/followers", rt.followUser)
	rt.router.DELETE("/users/profile/followers", rt.unfollowUser)
	rt.router.POST("/users/profile/post", rt.uploadPhoto)
	rt.router.DELETE("/users/profile/post", rt.removePost)
	rt.router.POST("/post/like", rt.likePhoto)       //passes photo id in the query params
	rt.router.POST("/users/profile/ban", rt.banUser) //username passed in the query
	rt.router.DELETE("/users/profile/ban", rt.unBanUser)
	rt.router.DELETE("/post/like", rt.unLikePhoto)
	rt.router.POST("/post/comment", rt.addComment)
	rt.router.DELETE("/post/comment", rt.removeComment)
	rt.router.GET("/users/profile/feed", rt.getFeed)
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
