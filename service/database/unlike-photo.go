package database

func (db *appdbimpl) RemoveLike(postId string, likerId string) error {

	query := `
		DELETE FROM likes 
		WHERE user_id = ? AND post_id = ?
    `

	_, err := db.c.Exec(query, likerId, postId)
	return err
}
