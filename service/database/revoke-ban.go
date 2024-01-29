package database

import (
	"myproject/service/types"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) RevokeBan(profile types.User, viewerId string) error {
	query := `DELETE FROM bans 
	WHERE user_id = ? AND banned_id = ?`

	// Execute the DELETE query
	_, err := db.c.Exec(query, viewerId, profile.Id)
	return err

}
