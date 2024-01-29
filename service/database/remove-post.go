package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) RemovePost(ownerId string, postId string) error {
	query := `DELETE FROM posts
	WHERE post_id = ? and owner_id = ?;`

	// Execute the DELETE query
	_, err := db.c.Exec(query, postId, ownerId)
	return err

}
