package api

import "time"

type User struct {
	Id        string
	Username  string
	Post      []Post
	Follower  []Follower
	Following []Follower
	Banned    []Follower
}

type Post struct {
	PostId   string
	OwnerId  string
	photo    string
	Likes    []uint64
	Comments []Comment
	time     time.Time
}
type Comment struct {
	User    *User
	message string
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
