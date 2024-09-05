package database

func (db *appdbimpl) CheckBan(toCheckUserId string, userId string) error {
	var placeholder string
	err := db.c.QueryRow("SELECT userId FROM Bans WHERE userId=? AND bannedUserId=?", userId, toCheckUserId).Scan(&placeholder)
	return err
}

func (db *appdbimpl) BanUser(userId string, toBanUserId string) error {

	_, err := db.c.Exec(`INSERT INTO Bans (userId, bannedUserId) VALUES (?, ?)`, userId, toBanUserId)
	return err
}

func (db *appdbimpl) DeleteBan(userId string, toUnbanUserId string) error {
	_, err := db.c.Exec(`DELETE FROM Bans WHERE userId=? AND bannedUserId=?`, userId, toUnbanUserId)
	return err
}
