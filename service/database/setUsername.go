package database

import "myproject/service/types"

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetUsername(user types.User, newUsername string) error {
	_, err := db.c.Exec("UPDATE users SET (name) VALUES (%s) WHERE id = %s", newUsername, user.Id)
	return err
}
