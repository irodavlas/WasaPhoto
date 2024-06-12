package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.doLogin)
	rt.router.PUT("/settings/username", rt.setNewUsername)

	rt.router.POST("/posts/", rt.uploadPhoto)
	rt.router.DELETE("/posts/:postId", rt.removePost)

	rt.router.PUT("/posts/:postId/like", rt.likePhoto)
	rt.router.DELETE("/posts/:postId/like", rt.unLikePhoto)

	rt.router.POST("/posts/:postId/comment", rt.addComment)
	rt.router.DELETE("/posts/:postId/comments/:commentId", rt.removeComment)

	rt.router.PUT("/follow/:userId", rt.followUser)
	rt.router.DELETE("/follow/:userId", rt.unfollowUser)

	rt.router.PUT("/ban/:username", rt.banUser)
	rt.router.DELETE("/ban/:username", rt.unBanUser)

	// get
	rt.router.GET("/users/:userId/profile", rt.getProfile)
	rt.router.GET("/profile/feed", rt.getFeed)

	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
