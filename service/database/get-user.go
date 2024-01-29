package database

import "myproject/service/types"

func (db *appdbimpl) GetUser(id string, username string) (types.User, error) {
	var user types.User
	err := db.c.QueryRow("SELECT * FROM users WHERE id = ? or username = ?", id, username).
		Scan(&user.Id, &user.Username)

	return user, err
}
