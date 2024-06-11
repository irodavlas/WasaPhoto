package database

func (db *appdbimpl) UpdateUsername(id string, newUsername string) error {
	_, err := db.c.Exec(`
	UPDATE users
	SET username = ?
	WHERE id = ?
    `, newUsername, id)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(`
	UPDATE posts
	SET owner_username = ?
	WHERE posts.owner_id = ?
    `, newUsername, id)
	if err != nil {
		return err
	}
	return err
}
