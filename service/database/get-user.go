package database

func (db *appdbimpl) GetUser() (string, error) {
	var name string
	err := db.c.QueryRow("SELECT * FROM users WHERE name=user_01").Scan(&name)
	return name, err
}
