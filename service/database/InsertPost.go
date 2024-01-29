package database

import (
	"myproject/service/types"
	"time"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) InsertPost(post types.Post) error {
	query := ` INSERT INTO posts (post_id, owner_id, owner_username, photo, created_at)
				VALUES (?, ?, ?, ?, ?)`

	_, err := db.c.Exec(query, post.PostId, post.OwnerId, post.OwnerUsername, post.Photo, time.Now())
	return err
}
