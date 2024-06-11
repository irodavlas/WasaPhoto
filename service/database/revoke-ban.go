package database

import (
	"myproject/service/types"
)

func (db *appdbimpl) RevokeBan(profile types.User, viewerId string) error {
	query := `DELETE FROM bans 
	WHERE user_id = ? AND banned_id = ?`

	_, err := db.c.Exec(query, viewerId, profile.Id)
	return err

}
