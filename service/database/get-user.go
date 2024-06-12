package database

import "myproject/service/types"

func (db *appdbimpl) GetUser(id string, username string) (types.User, error) {
	var user types.User

	err := db.c.QueryRow("SELECT * FROM users WHERE id = ? or username = ?", id, username).
		Scan(&user.Id, &user.Username)

	return user, err
}
func (db *appdbimpl) GetUsers() ([]string, error) {
	var users []string
	query := `
	SELECT username
	FROM users;`

	rows, err := db.c.Query(query)
	if err != nil || rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}

		users = append(users, username)
	}

	return users, err
}
