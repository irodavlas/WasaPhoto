package database

import (
	"myproject/service/types"
)

func (db *appdbimpl) InsertLike(likeid string, postId string, user types.User) error {

	query := `
	INSERT INTO likes (like_id, user_id, post_id)
	SELECT ?, ?, ?
	WHERE EXISTS (
		SELECT post_id 
		FROM posts 
		WHERE post_id = ?  -- Check if the post_id exists in the posts table
	) AND NOT EXISTS (
		SELECT 1
		FROM likes
		WHERE user_id = ? AND post_id = ?  -- Check if there is no existing row with the same user_id and post_id in the likes table
	)
    `

	_, err := db.c.Exec(query, likeid, user.Id, postId, postId, user.Id, postId)
	return err
}
