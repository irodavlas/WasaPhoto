package database

import "myproject/service/types"

func (db *appdbimpl) GetUser(id string, username string) (types.User, error) {
	var user types.User
	err := db.c.QueryRow("SELECT * FROM users WHERE username = %s or id = %s", username, id).Scan(&user.Id, &user.Username)

	return user, err
}
