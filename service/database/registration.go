package database

func (db *appdbimpl) RegisterUser(username string, userId string) error {

	_, err := db.c.Exec(`INSERT INTO Users (userId, username) VALUES (?, ?)`, userId, username)
	return err
}
