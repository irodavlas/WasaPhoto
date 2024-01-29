package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) RemoveFollow(followerId string, userId string) error {
	query := `DELETE FROM followers 
	WHERE follower_id = ? AND user_id = ?`

	// Execute the DELETE query
	_, err := db.c.Exec(query, followerId, userId)
	return err

}
