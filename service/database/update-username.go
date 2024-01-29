package database

func (db *appdbimpl) UpdateUsername(id string, newUsername string) error {
	_, err := db.c.Exec("UPDATE users SET username = ? WHERE id = ?", newUsername, id)
	return err

}
