package database

func (db *appdbimpl) FollowUser(userId string, followingUserId string) error {

	_, err := db.c.Exec(`INSERT INTO Followers (userId, followingUserId) VALUES (?, ?)`, userId, followingUserId)
	return err
}

func (db *appdbimpl) CheckFollow(userId string, followingUserId string) error {
	var placeholder string
	err := db.c.QueryRow("SELECT userId FROM Followers WHERE userId=? AND followingUserId=?", userId, followingUserId).Scan(&placeholder)
	return err
}

func (db *appdbimpl) DeleteFollow(userId string, toUnfollowUserId string) error {
	_, err := db.c.Exec(`DELETE FROM Followers WHERE userId=? AND followingUserId=?`, userId, toUnfollowUserId)
	return err
}

func (db *appdbimpl) RetrieveFollowing(userId string) (int, error) {
	var followingCount int
	err := db.c.QueryRow(`SELECT COUNT(userId) AS likeCount FROM Followers WHERE userId=?;`, userId).Scan(&followingCount)
	return followingCount, err
}

func (db *appdbimpl) RetrieveFollowers(userId string) (int, error) {
	var followersCount int
	err := db.c.QueryRow(`SELECT COUNT(userId) AS likeCount FROM Followers WHERE followingUserId=?;`, userId).Scan(&followersCount)
	return followersCount, err
}
