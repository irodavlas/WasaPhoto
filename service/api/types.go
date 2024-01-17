package api

type User struct {
	Username string
	Id       string
}

type Photo struct {
	PhotoID  string
	OwnerID  string
	Img      string
	UpDate   string
	Likes    []*User
	Comments []*Comment
}
type Comment struct {
	User    *User
	message string
}
type UserProfile struct {
	User      *User
	Post      map[string]*Photo
	Follower  map[string]*User
	Following map[string]*User
	Banned    map[string]*User
}
