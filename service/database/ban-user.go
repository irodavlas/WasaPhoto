package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) BanUser(bannedUserId string, userId string) error {
	err := db.RemoveFollow(bannedUserId, userId)
	if err != nil {
		return err
	}
	err = db.RemoveFollow(userId, bannedUserId)
	if err != nil {
		return err
	}
	query := "INSERT INTO bans (banned_id, user_id) VALUES (?, ?)"
	_, err = db.c.Exec(query, bannedUserId, userId)
	return err
}
