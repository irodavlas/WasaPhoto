package database

func (db *appdbimpl) RetrieveId(username string) (string, error) {
	var userId string
	err := db.c.QueryRow(`SELECT userId FROM Users WHERE username=?`, username).Scan(&userId)
	return userId, err
}

func (db *appdbimpl) RetrieveUsername(userId string) (string, error) {
	var username string
	err := db.c.QueryRow(`SELECT username FROM Users WHERE userId=?`, userId).Scan(&username)
	return username, err
}
