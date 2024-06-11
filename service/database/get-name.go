package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetName() (string, error) {
	var name string
	var id string
	err := db.c.QueryRow("SELECT * FROM users").Scan(&name, &id)
	return name, err
}
