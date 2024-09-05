package api

import "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"

var streamDatabase = make(map[string][]database.Photo)

const notExist = ` doesn't exist`

type Photo struct {
	Image    []byte    `json:"image"`
	PhotoId  string    `json:"photoId"`
	Username string    `json:"username"`
	Comments []Comment `json:"comments"`
	Likes    int       `json:"likes"`
	LikeId   string    `json:"likeId"`
}

type Comment struct {
	Content   string `json:"content"`
	CommentId string `json:"commentId"`
	Username  string `json:"username"`
}

type Profile struct {
	Followers int              `json:"followers"`
	Following int              `json:"following"`
	Photos    []database.Photo `json:"photos"`
	Banned    bool             `json:"banned"`
}
