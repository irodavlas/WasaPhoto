package api

type User struct {
	Username string
	Id       string
}
type Params struct {
	Username string
}

type Photo struct {
	PhotoID int
	OwnerID string
	Img     string
	UpDate  string
}

type UserProfile struct {
	User      User
	Post      []Photo
	Follower  []User
	Following []User
}
