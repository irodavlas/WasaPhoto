package database

import "myproject/service/types"

func (db *appdbimpl) InsertUser(user types.User) error {
	_, err := db.c.Exec("INSERT INTO users (id, username) VALUES (?, ?)", user.Id, user.Username)
	return err
}
