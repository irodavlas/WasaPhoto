package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) RemoveLike(postId string, likerId string) error {

	query := `
		DELETE FROM likes 
		WHERE user_id = ? AND post_id = ?
    `

	_, err := db.c.Exec(query, likerId, postId)
	return err
}
