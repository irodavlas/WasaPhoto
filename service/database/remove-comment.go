package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) RemoveComment(postId string, ownerId string, commentId string) error {
	query := `DELETE FROM comments 
	WHERE user_id = ? AND post_id = ? and comment_id = ?`

	// Execute the DELETE query
	_, err := db.c.Exec(query, ownerId, postId, commentId)
	return err

}
