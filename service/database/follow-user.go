package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) InsertFollower(userId string, followerId string) error {
	_, err := db.c.Exec("INSERT INTO followers (user_id, follower_id) VALUES (?, ?)", userId, followerId)
	return err
}
