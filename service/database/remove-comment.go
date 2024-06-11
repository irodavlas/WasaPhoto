package database

func (db *appdbimpl) RemoveComment(postId string, ownerId string, commentId string) error {
	query := `DELETE FROM comments 
	WHERE user_id = ? AND post_id = ? and comment_id = ?`

	_, err := db.c.Exec(query, ownerId, postId, commentId)
	return err

}
