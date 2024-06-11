package database

func (db *appdbimpl) InsertFollower(userId string, followerId string) error {

	_, err := db.c.Exec("INSERT OR IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)", userId, followerId)

	return err
}
