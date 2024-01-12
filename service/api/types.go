package api

type User struct {
	Username string
	Id       string
}

type Photo struct {
	PhotoID int
	OwnerID string
	Img     string
	UpDate  string
}

type UserProfile struct {
	User      *User
	Post      []Photo
	Follower  map[string]*User
	Following map[string]*User
}
