package database

import (
	"myproject/service/api"
)

func (db *appdbimpl) GetUser(username string) (api.User, error) {
	var user api.User
	err := db.c.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.Id, &user.Username)

	return user, err
}
