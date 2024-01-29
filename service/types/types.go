package types

import "time"

type User struct {
	Id        string
	Username  string
	Npost     int
	Post      []Post
	Follower  []Follower
	Following []Follower
	Banned    []Follower
}

type Post struct {
	PostId        string
	OwnerId       string
	OwnerUsername string
	Photo         string
	Likes         []string
	Comments      []Comment
	Time          string
}

type Feed struct {
	Feed []Post
}
type Comment struct {
	CommentId string
	Username  string
	Message   string
}

type Follower struct {
	Id       string
	Username string
}

func NewUser(id string, username string) *User {
	return &User{
		Id:       id,
		Username: username,
	}
}
func NewPost(postId string, ownerId string, photo string, ownerUsername string) *Post {
	return &Post{
		PostId:        postId,
		OwnerId:       ownerId,
		OwnerUsername: ownerUsername,
		Photo:         photo,
		Time:          time.Now().Format("2006-01-02 15:04:05"),
	}
}
