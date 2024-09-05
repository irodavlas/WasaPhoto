package database

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
