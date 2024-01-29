package database

import "myproject/service/types"

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) InsertUser(user types.User) error {
	_, err := db.c.Exec("INSERT INTO users (id, username) VALUES (?, ?)", user.Id, user.Username)
	return err
}
