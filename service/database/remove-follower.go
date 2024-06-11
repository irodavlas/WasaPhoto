package database

func (db *appdbimpl) RemoveFollow(followerId string, userId string) error {
	query := `DELETE FROM followers 
	WHERE follower_id = ? AND user_id = ?`

	_, err := db.c.Exec(query, followerId, userId)
	return err

}
