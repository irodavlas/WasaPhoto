package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetName(userId string, username string) error {
	_, err := db.c.Exec("UPDATE Users SET username=? WHERE userId=?", username, userId)
	return err
}
