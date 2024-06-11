package database

func (db *appdbimpl) RemovePost(ownerId string, postId string) error {
	query := `DELETE FROM posts
	WHERE post_id = ? and owner_id = ?;`

	_, err := db.c.Exec(query, postId, ownerId)
	return err

}
