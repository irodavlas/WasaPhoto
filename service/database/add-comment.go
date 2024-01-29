package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) InsertComment(commentId string, postId string, viewerId string, content string) error {

	query := `
        INSERT INTO comments (comment_id, user_id, post_id, comment_content)
        SELECT ?, ?, ?, ?
        WHERE EXISTS (SELECT post_id FROM posts WHERE post_id = ?)
    `

	_, err := db.c.Exec(query, commentId, viewerId, postId, content, postId)
	return err
}
